package controller

import (
	"errors"
	"github.com/UPrefer/StorageService/mocks"
	"github.com/UPrefer/StorageService/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type ArtifactReadTestSuite struct {
	suite.Suite

	artifactId   string
	artifactName string
	uploadDate   time.Time

	context               *gin.Context
	httpRecorder          *httptest.ResponseRecorder
	mockedArtifactService *mocks.IArtifactService
	artifactController    *ArtifactController
}

func (suite *ArtifactReadTestSuite) SetupTest() {
	suite.artifactId = "artifactId1"
	suite.artifactName = "artifactName1"
	suite.uploadDate = time.Date(2015, time.August, 15, 20, 00, 00, 00, time.UTC)

	suite.mockedArtifactService = &mocks.IArtifactService{}
	suite.artifactController = &ArtifactController{artifactService: suite.mockedArtifactService}

	suite.httpRecorder = httptest.NewRecorder()
	suite.context, _ = gin.CreateTestContext(suite.httpRecorder)
	suite.context.Params = []gin.Param{{"artifact_id", suite.artifactId}}
}

func TestArtifactRead(t *testing.T) {
	suite.Run(t, new(ArtifactReadTestSuite))
}

func (suite *ArtifactReadTestSuite) Test_ShouldReturnHTTP200_AndId_WhenArtifactWaitsForUpload() {
	//GIVEN
	var (
		expectedStatus = http.StatusOK
		expectedBody   = "{\"uuid\":\"artifactId1\",\"name\":\"artifactName1\"}"
	)

	suite.mockedArtifactService.On("ReadArtifact", suite.artifactId).Return(&model.ArtifactDTO{Uuid: suite.artifactId, Name: suite.artifactName}, nil)

	//WHEN
	suite.artifactController.Get(suite.context)

	//THEN
	suite.Equal(expectedStatus, suite.context.Writer.Status())
	suite.Equal(expectedBody, suite.httpRecorder.Body.String())
}

func (suite *ArtifactReadTestSuite) Test_ShouldReturnHTTP200_AndId_WhenArtifactUploaded() {
	//GIVEN
	var (
		expectedStatus = http.StatusOK
		expectedBody   = "{\"uuid\":\"artifactId1\",\"name\":\"artifactName1\",\"length\":12,\"contentType\":\"application/data\",\"md5sum\":\"anMd5sum\",\"uploadDate\":\"2015-08-15T20:00:00Z\"}"
	)

	suite.mockedArtifactService.On("ReadArtifact", suite.artifactId).Return(&model.ArtifactDTO{Uuid: suite.artifactId, Name: suite.artifactName, ContentType: "application/data", Length: 12, Md5: "anMd5sum", UploadDate: &suite.uploadDate}, nil)

	//WHEN
	suite.artifactController.Get(suite.context)

	//THEN
	suite.Equal(expectedStatus, suite.context.Writer.Status())
	suite.Equal(expectedBody, suite.httpRecorder.Body.String())
}

func (suite *ArtifactReadTestSuite) Test_ShouldReturnHTTP404_AndEmptyBody_WhenArtifactDoesNotExist() {
	//GIVEN
	var (
		expectedStatus = http.StatusNotFound
		expectedBody   = ""
	)

	suite.mockedArtifactService.On("ReadArtifact", suite.artifactId).Return(nil, nil)

	//WHEN
	suite.artifactController.Get(suite.context)

	//THEN
	suite.Equal(expectedStatus, suite.context.Writer.Status())
	suite.Equal(expectedBody, suite.httpRecorder.Body.String())
}

func (suite *ArtifactReadTestSuite) Test_ShouldReturHTTP500_AndEmptyBody_WhenGetArtifactFails() {
	//GIVEN
	var (
		expectedStatus = http.StatusInternalServerError
		expectedBody   = ""
	)

	suite.mockedArtifactService.On("ReadArtifact", suite.artifactId).Return(nil, errors.New("random error"))

	//WHEN
	suite.artifactController.Get(suite.context)

	//THEN
	suite.Equal(expectedStatus, suite.context.Writer.Status())
	suite.Equal(expectedBody, suite.httpRecorder.Body.String())
}
