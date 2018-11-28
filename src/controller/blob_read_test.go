package controller

import (
	"bytes"
	"errors"
	"github.com/UPrefer/StorageService/mocks"
	"github.com/UPrefer/StorageService/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

type BlobReadTestSuite struct {
	suite.Suite

	context      *gin.Context
	httpRecorder *httptest.ResponseRecorder

	blobService    *mocks.IBlobService
	blobController *BlobController

	randomReadCloser io.ReadCloser

	content       []byte
	contentType   string
	contentLength int64
	artifactId    string
}

func (suite *BlobReadTestSuite) SetupTest() {
	suite.artifactId = "artifact1"
	suite.contentType = "application/data"
	suite.content = []byte("pouet")
	suite.contentLength = int64(len(suite.content))

	suite.randomReadCloser = ioutil.NopCloser(bytes.NewReader(suite.content))

	suite.httpRecorder = httptest.NewRecorder()
	suite.context, _ = gin.CreateTestContext(suite.httpRecorder)
	suite.context.Request = &http.Request{}
	suite.context.Params = []gin.Param{{"artifact_id", suite.artifactId}}

	suite.blobService = &mocks.IBlobService{}

	suite.blobController = &BlobController{blobService: suite.blobService}
}

func TestBlobRead(t *testing.T) {
	suite.Run(t, new(BlobReadTestSuite))
}

func (suite *BlobReadTestSuite) Test_ShouldReturnHttp200_AndFile_AndContentType_WhenNoError() {
	//GIVEN
	var (
		expectedStatus = http.StatusOK
	)

	suite.blobService.On("ReadBlob", suite.artifactId).Return(suite.contentType, suite.contentLength, suite.randomReadCloser, nil)

	//WHEN
	suite.blobController.Get(suite.context)

	//THEN
	suite.Equal(expectedStatus, suite.context.Writer.Status())
	suite.Equal(suite.content, suite.httpRecorder.Body.Bytes())
	suite.Equal(suite.contentType, suite.httpRecorder.Header().Get("Content-Type"))
	suite.Equal(strconv.Itoa(int(suite.contentLength)), suite.httpRecorder.Header().Get("Content-Length"))
}

func (suite *BlobReadTestSuite) Test_ShouldReturnHttp404_AndEmptyBody_WhenArtifactNotUploaded() {
	//GIVEN
	var (
		expectedStatus = http.StatusNotFound
		expectedBody   = ""
	)

	suite.blobService.On("ReadBlob", suite.artifactId).Return("", int64(0), nil, service.ErrArtifactNotFound)

	//WHEN
	suite.blobController.Get(suite.context)

	//THEN
	suite.Equal(expectedStatus, suite.context.Writer.Status())
	suite.Equal(expectedBody, suite.httpRecorder.Body.String())
}

func (suite *BlobReadTestSuite) Test_ShouldReturnHttp500_AndEmptyBody_WhenTreatmentFails() {
	//GIVEN
	var (
		expectedStatus = http.StatusInternalServerError
		expectedBody   = ""
	)

	suite.blobService.On("ReadBlob", suite.artifactId).Return("", int64(0), nil, errors.New("random error"))

	//WHEN
	suite.blobController.Get(suite.context)

	//THEN
	suite.Equal(expectedStatus, suite.context.Writer.Status())
	suite.Equal(expectedBody, suite.httpRecorder.Body.String())
}
