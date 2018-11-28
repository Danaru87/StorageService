package service

import (
	"errors"
	"github.com/UPrefer/StorageService/mocks"
	"github.com/UPrefer/StorageService/model"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ArtifactCreationTestSuite struct {
	suite.Suite
	mockedUtilsService *mocks.IUtilsService
	mockedArtifactDao  *mocks.IArtifactDao
	artifactService    *ArtifactService
}

func (suite *ArtifactCreationTestSuite) SetupTest() {
	suite.mockedUtilsService = &mocks.IUtilsService{}
	suite.mockedArtifactDao = &mocks.IArtifactDao{}
	suite.artifactService = NewArtifactService(suite.mockedUtilsService, suite.mockedArtifactDao)
}

func TestArtifactCreation(t *testing.T) {
	suite.Run(t, new(ArtifactCreationTestSuite))
}

func (suite *ArtifactCreationTestSuite) Test_ShouldReturnCreatedArtifact_WhenNoError() {
	//GIVEN
	var (
		artifactId                = "artifact1"
		expectedArtifactDto       = model.ArtifactDTO{Uuid: artifactId}
		expectedError       error = nil
	)

	suite.mockedUtilsService.On("NewUUID").Return(artifactId, nil)
	suite.mockedArtifactDao.On("CreateArtifact", &expectedArtifactDto).Return(nil)

	//WHEN
	var createdArtifactDto, err = suite.artifactService.CreateArtifact(&model.ArtifactDTO{})

	//THEN
	suite.Assert().Equal(expectedError, err)
	suite.Assert().Equal(&expectedArtifactDto, createdArtifactDto)
}

func (suite *ArtifactCreationTestSuite) Test_ShouldReturnEncounteredError_AndAttemptedToCreateDto_WhenDaoFails() {
	//GIVEN
	var (
		artifactId          = "artifact1"
		expectedArtifactDto = &model.ArtifactDTO{Uuid: artifactId}
		expectedError       = errors.New("Any Dao Error")
	)

	suite.mockedUtilsService.On("NewUUID").Return(artifactId, nil)
	suite.mockedArtifactDao.On("CreateArtifact", expectedArtifactDto).Return(expectedError)

	//WHEN
	var createdArtifactDto, err = suite.artifactService.CreateArtifact(&model.ArtifactDTO{})

	//THEN
	suite.Assert().Equal(expectedError, err)
	suite.Assert().Equal(expectedArtifactDto, createdArtifactDto)
}

func (suite *ArtifactCreationTestSuite) Test_ShouldReturnEncounteredError_WhenUUIDGenerationFails() {
	//GIVEN
	var (
		expectedArtifactDto *model.ArtifactDTO = nil
		expectedError                          = errors.New("Any UUID generation Error")
	)

	suite.mockedUtilsService.On("NewUUID").Return("", expectedError)

	//WHEN
	var createdArtifactDto, err = suite.artifactService.CreateArtifact(&model.ArtifactDTO{})

	//THEN
	suite.Assert().Equal(expectedError, err)
	suite.Assert().Equal(expectedArtifactDto, createdArtifactDto)
}
