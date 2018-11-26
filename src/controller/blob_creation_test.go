package controller

import (
	"errors"
	"github.com/UPrefer/StorageService/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type BlobCreationTestSuite struct {
	suite.Suite

	context           *gin.Context
	httpRecorder      *httptest.ResponseRecorder
	blobController    *BlobController
	mockedBlobService *service.MockedBlobService
}

func (suite *BlobCreationTestSuite) SetupTest() {
	suite.httpRecorder = httptest.NewRecorder()
	suite.context, _ = gin.CreateTestContext(suite.httpRecorder)
	suite.context.Request = &http.Request{}

	suite.mockedBlobService = &service.MockedBlobService{}

	suite.blobController = &BlobController{blobService: suite.mockedBlobService}
}

func TestBlobCreation(t *testing.T) {
	suite.Run(t, new(BlobCreationTestSuite))
}

func (suite *BlobCreationTestSuite) Test_ShouldReturnHttp404_AndEmptyBody_WhenArtifactDoesNotExist() {
	//GIVEN
	var (
		expectedStatus = http.StatusNotFound
		expectedBody   = ""
	)

	suite.mockedBlobService.ExpectedSaveBlobError = service.ErrArtifactNofFound

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

	suite.mockedBlobService.ExpectedSaveBlobError = errors.New("random technical error")

	//WHEN
	suite.blobController.Put(suite.context)

	//THEN
	suite.Equal(expectedStatus, suite.context.Writer.Status())
	suite.Equal(expectedBody, suite.httpRecorder.Body.String())
}

func (suite *BlobCreationTestSuite) Test_ShouldReturnHttp200_AndEmptyBody_WhenUploadOk() {
	//GIVEN
	var (
		expectedStatus = http.StatusOK
		expectedBody   = ""
	)

	//WHEN
	suite.blobController.Put(suite.context)

	//THEN
	suite.Equal(expectedStatus, suite.context.Writer.Status())
	suite.Equal(expectedBody, suite.httpRecorder.Body.String())
}
