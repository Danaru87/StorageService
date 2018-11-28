package controller

import (
	"bytes"
	"errors"
	"github.com/UPrefer/StorageService/mocks"
	"github.com/UPrefer/StorageService/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ArtifactCreationTestSuite struct {
	suite.Suite
	artifactName        string
	validArtifactJson   *bytes.Buffer
	invalidArtifactJson *bytes.Buffer

	context               *gin.Context
	httpRecorder          *httptest.ResponseRecorder
	mockedArtifactService *mocks.IArtifactService
	artifactController    *ArtifactController
}

func (suite *ArtifactCreationTestSuite) SetupTest() {
	suite.mockedArtifactService = &mocks.IArtifactService{}
	suite.artifactController = NewArtifactController(suite.mockedArtifactService)
	suite.httpRecorder = httptest.NewRecorder()
	suite.context, _ = gin.CreateTestContext(suite.httpRecorder)

	suite.artifactName = "artifact1"
	suite.validArtifactJson = bytes.NewBufferString("{\"name\": \"artifact1\"}")
	suite.invalidArtifactJson = bytes.NewBufferString("{}")
}

func TestArtifactCreation(t *testing.T) {
	suite.Run(t, new(ArtifactCreationTestSuite))
}

func (suite *ArtifactCreationTestSuite) Test_ShouldSetStatusHttp201_AndEmptyBody_AndLocationHeader_WhenCreationOk() {
	//GIVEN
	var (
		expectedStatus         = http.StatusCreated
		expectedBody           = ""
		expectedLocationHeader = "/artifact/0000-1111-2222-3333"
	)

	suite.mockedArtifactService.On("CreateArtifact", &model.ArtifactDTO{Name: suite.artifactName}).Return(&model.ArtifactDTO{Uuid: "0000-1111-2222-3333", Name: suite.artifactName}, nil)

	suite.context.Request = httptest.NewRequest(http.MethodPost, "/artifact", suite.validArtifactJson)

	//WHEN
	suite.artifactController.Post(suite.context)

	//THEN
	suite.Equal(expectedStatus, suite.context.Writer.Status())
	suite.Equal(expectedBody, suite.httpRecorder.Body.String())
	suite.Equal(expectedLocationHeader, suite.httpRecorder.Header().Get("Location"))
}

func (suite *ArtifactCreationTestSuite) Test_ShouldSetStatusHttp500_AndEmptyBody_WhenCreationFails() {
	//GIVEN
	var (
		expectedStatus = http.StatusInternalServerError
		expectedBody   = ""
	)

	suite.mockedArtifactService.On("CreateArtifact", &model.ArtifactDTO{Name: suite.artifactName}).Return(nil, errors.New("random failure"))

	suite.context.Request = httptest.NewRequest(http.MethodPost, "/artifact", suite.validArtifactJson)

	//WHEN
	suite.artifactController.Post(suite.context)

	//THEN
	suite.Equal(expectedStatus, suite.context.Writer.Status())
	suite.Equal(expectedBody, suite.httpRecorder.Body.String())
}

func (suite *ArtifactCreationTestSuite) Test_ShouldSetHttpStatus400_AndEmptyBody_WhenDataInputIsInvalid() {
	//GIVEN
	var (
		expectedStatus = http.StatusBadRequest
		expectedBody   = ""
	)

	suite.context.Request = httptest.NewRequest(http.MethodPost, "/artifact", suite.invalidArtifactJson)

	//WHEN
	suite.artifactController.Post(suite.context)

	//THEN
	suite.Equal(expectedStatus, suite.context.Writer.Status())
	suite.Equal(expectedBody, suite.httpRecorder.Body.String())
}
