package controller

import (
	"bytes"
	"errors"
	"github.com/UPrefer/StorageService/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockedArtifactService struct {
	mock.Mock
	expectedUuid  string
	expectedError error
}

func (artifactService *MockedArtifactService) CreateArtifact(artifactDto *model.ArtifactDTO) (*model.ArtifactDTO, error) {
	artifactDto.Uuid = artifactService.expectedUuid
	return artifactDto, artifactService.expectedError
}

type ArtifactCreationTestSuite struct {
	suite.Suite
	validArtifactJson   *bytes.Buffer
	invalidArtifactJson *bytes.Buffer

	context               *gin.Context
	httpRecorder          *httptest.ResponseRecorder
	mockedArtifactService *MockedArtifactService
	artifactController    *ArtifactController
}

func (suite *ArtifactCreationTestSuite) SetupTest() {
	suite.mockedArtifactService = new(MockedArtifactService)
	suite.artifactController = NewArtifactController(suite.mockedArtifactService)
	suite.httpRecorder = httptest.NewRecorder()
	suite.context, _ = gin.CreateTestContext(suite.httpRecorder)

	suite.validArtifactJson = bytes.NewBufferString("{\"name\": \"artifactName\"}")
	suite.invalidArtifactJson = bytes.NewBufferString("{}")
}

func TestArtifactCreation(t *testing.T) {
	suite.Run(t, new(ArtifactCreationTestSuite))
}

func (suite *ArtifactCreationTestSuite) Test_ShouldSetStatusHttp201_AndEmptyBody_AndLocationHeader_WhenCreationOk() {
	//GIVEN
	var expectedStatus = http.StatusCreated
	var expectedBody = ""
	var expectedLocationHeader = "/artifact/0000-1111-2222-3333"

	suite.mockedArtifactService.expectedUuid = "0000-1111-2222-3333"

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
	var expectedStatus = http.StatusInternalServerError
	var expectedBody = ""

	suite.mockedArtifactService.expectedError = errors.New("random failure")

	suite.context.Request = httptest.NewRequest(http.MethodPost, "/artifact", suite.validArtifactJson)

	//WHEN
	suite.artifactController.Post(suite.context)

	//THEN
	suite.Equal(expectedStatus, suite.context.Writer.Status())
	suite.Equal(expectedBody, suite.httpRecorder.Body.String())
}

func (suite *ArtifactCreationTestSuite) Test_ShouldSetHttpStatus400_AndEmptyBody_WhenDataInputIsInvalid() {
	//GIVEN
	var expectedStatus = http.StatusBadRequest
	var expectedBody = ""

	suite.context.Request = httptest.NewRequest(http.MethodPost, "/artifact", suite.invalidArtifactJson)

	//WHEN
	suite.artifactController.Post(suite.context)

	//THEN
	suite.Equal(expectedStatus, suite.context.Writer.Status())
	suite.Equal(expectedBody, suite.httpRecorder.Body.String())
}
