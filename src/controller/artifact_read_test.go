package controller

import (
	"errors"
	"github.com/UPrefer/StorageService/model"
	"github.com/UPrefer/StorageService/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ArtifactReadTestSuite struct {
	suite.Suite

	context               *gin.Context
	httpRecorder          *httptest.ResponseRecorder
	mockedArtifactService *service.MockedArtifactService
	artifactController    *ArtifactController
}

func (suite *ArtifactReadTestSuite) SetupTest() {
	suite.mockedArtifactService = &service.MockedArtifactService{}
	suite.artifactController = &ArtifactController{artifactService: suite.mockedArtifactService}

	suite.httpRecorder = httptest.NewRecorder()
	suite.context, _ = gin.CreateTestContext(suite.httpRecorder)
}

func TestArtifactRead(t *testing.T) {
	suite.Run(t, new(ArtifactReadTestSuite))
}

func (suite *ArtifactReadTestSuite) Test_ShouldReturnHTTP200_AndId_WhenArtifactWaitsForUpload() {
	//GIVEN
	var expectedStatus = http.StatusOK
	var expectedBody = "{\"uuid\":\"a1\"}"

	suite.mockedArtifactService.ReadArtifact_ExpectedArtifact = &model.ArtifactDTO{Uuid: "a1"}
	suite.mockedArtifactService.ReadArtifact_ExpectedError = nil

	suite.context.Params = []gin.Param{{Key: "artifact_id", Value: "a1"}}

	//WHEN
	suite.artifactController.Get(suite.context)

	//THEN
	suite.Equal(expectedStatus, suite.context.Writer.Status())
	suite.Equal(expectedBody, suite.httpRecorder.Body.String())
}

func (suite *ArtifactReadTestSuite) Test_ShouldReturnHTTP404_AndEmptyBody_WhenArtifactDoesNotExist() {
	//GIVEN
	var expectedStatus = http.StatusNotFound
	var expectedBody = ""

	suite.mockedArtifactService.ReadArtifact_ExpectedError = nil
	suite.mockedArtifactService.ReadArtifact_ExpectedArtifact = nil

	//WHEN
	suite.artifactController.Get(suite.context)

	//THEN
	suite.Equal(expectedStatus, suite.context.Writer.Status())
	suite.Equal(expectedBody, suite.httpRecorder.Body.String())
}

func (suite *ArtifactReadTestSuite) Test_ShouldReturHTTP500_AndEmptyBody_WhenGetArtifactFails() {
	//GIVEN
	var expectedStatus = http.StatusInternalServerError
	var expectedBody = ""

	suite.mockedArtifactService.ReadArtifact_ExpectedArtifact = nil
	suite.mockedArtifactService.ReadArtifact_ExpectedError = errors.New("random error")

	//WHEN
	suite.artifactController.Get(suite.context)

	//THEN
	suite.Equal(expectedStatus, suite.context.Writer.Status())
	suite.Equal(expectedBody, suite.httpRecorder.Body.String())
}
