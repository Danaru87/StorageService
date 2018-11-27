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
)

type ArtifactCreationTestSuite struct {
	suite.Suite

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

	suite.mockedArtifactService.On("CreateArtifact").Return(&model.ArtifactDTO{Uuid: "0000-1111-2222-3333"}, nil)

	suite.context.Request = httptest.NewRequest(http.MethodPost, "/artifact", nil)

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

	suite.mockedArtifactService.On("CreateArtifact").Return(nil, errors.New("random failure"))

	suite.context.Request = httptest.NewRequest(http.MethodPost, "/artifact", nil)

	//WHEN
	suite.artifactController.Post(suite.context)

	//THEN
	suite.Equal(expectedStatus, suite.context.Writer.Status())
	suite.Equal(expectedBody, suite.httpRecorder.Body.String())
}
