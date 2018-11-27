package controller

import (
	"crypto/rand"
	"errors"
	"github.com/UPrefer/StorageService/mocks"
	"github.com/UPrefer/StorageService/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type BlobCreationTestSuite struct {
	suite.Suite

	context           *gin.Context
	httpRecorder      *httptest.ResponseRecorder
	blobController    *BlobController
	mockedBlobService *mocks.IBlobService

	randomReadCloser io.ReadCloser

	artifactId  string
	contentType string
}

func (suite *BlobCreationTestSuite) SetupTest() {
	suite.artifactId = "artifact1"
	suite.contentType = "application/data"

	suite.randomReadCloser = ioutil.NopCloser(io.LimitReader(rand.Reader, 255))

	suite.httpRecorder = httptest.NewRecorder()
	suite.context, _ = gin.CreateTestContext(suite.httpRecorder)
	suite.context.Request = &http.Request{}

	suite.context.Request.Body = suite.randomReadCloser
	suite.context.Params = []gin.Param{{"artifact_id", suite.artifactId}}
	suite.context.Request.Header = map[string][]string{"Content-Type": {suite.contentType}}

	suite.mockedBlobService = &mocks.IBlobService{}

	suite.blobController = &BlobController{blobService: suite.mockedBlobService}
}

func TestBlobCreation(t *testing.T) {
	suite.Run(t, new(BlobCreationTestSuite))
}

func (suite *BlobCreationTestSuite) Test_ShouldReturnHttp200_AndEmptyBody_WhenUploadOk() {
	//GIVEN
	var (
		expectedStatus = http.StatusOK
		expectedBody   = ""
	)

	suite.mockedBlobService.On("SaveBlob", suite.artifactId, suite.contentType, suite.randomReadCloser).Return(nil)

	//WHEN
	suite.blobController.Put(suite.context)

	//THEN
	suite.Equal(expectedStatus, suite.context.Writer.Status())
	suite.Equal(expectedBody, suite.httpRecorder.Body.String())
}

func (suite *BlobCreationTestSuite) Test_ShouldReturnHttp404_AndEmptyBody_WhenArtifactDoesNotExist() {
	//GIVEN
	var (
		expectedStatus = http.StatusNotFound
		expectedBody   = ""
	)

	suite.mockedBlobService.On("SaveBlob", suite.artifactId, suite.contentType, suite.randomReadCloser).Return(service.ErrArtifactNotFound)

	//WHEN
	suite.blobController.Put(suite.context)

	//THEN
	suite.Equal(expectedStatus, suite.context.Writer.Status())
	suite.Equal(expectedBody, suite.httpRecorder.Body.String())
}

func (suite *BlobCreationTestSuite) Test_ShouldReturnHttp409_AndEmptyBody_WhenArtifactAlreadyUploaded() {
	//GIVEN
	var (
		expectedStatus = http.StatusConflict
		expectedBody   = ""
	)

	suite.mockedBlobService.On("SaveBlob", suite.artifactId, suite.contentType, suite.randomReadCloser).Return(service.ErrArtifactAlreadyUploaded)

	//WHEN
	suite.blobController.Put(suite.context)

	//THEN
	suite.Equal(expectedStatus, suite.context.Writer.Status())
	suite.Equal(expectedBody, suite.httpRecorder.Body.String())
}

func (suite *BlobCreationTestSuite) Test_ShouldReturnHttp500_AndEmptyBody_WhenTechnicalErrorHappens() {
	//GIVEN
	var (
		expectedStatus = http.StatusInternalServerError
		expectedBody   = ""
	)

	suite.mockedBlobService.On("SaveBlob", suite.artifactId, suite.contentType, suite.randomReadCloser).Return(errors.New("random technical error"))

	//WHEN
	suite.blobController.Put(suite.context)

	//THEN
	suite.Equal(expectedStatus, suite.context.Writer.Status())
	suite.Equal(expectedBody, suite.httpRecorder.Body.String())
}
