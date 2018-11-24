package service

import (
	"errors"
	"github.com/UPrefer/StorageService/model"
	"github.com/stretchr/testify/suite"
	"testing"
)

type MockedArtifactDao struct {
	expectedError error
}

func (dao *MockedArtifactDao) CreateArtifact(*model.ArtifactDTO) error {
	return dao.expectedError
}

type ArtifactCreationTestSuite struct {
	suite.Suite
	mockedUtilsService *MockedUtilsService
	mockedArtifactDao  *MockedArtifactDao
	artifactService    *ArtifactService
}

func (suite *ArtifactCreationTestSuite) SetupTest() {
	suite.mockedUtilsService = &MockedUtilsService{}
	suite.mockedArtifactDao = &MockedArtifactDao{}
	suite.artifactService = NewArtifactService(suite.mockedUtilsService, suite.mockedArtifactDao)
}

func TestArtifactCreation(t *testing.T) {
	suite.Run(t, new(ArtifactCreationTestSuite))
}

func (suite *ArtifactCreationTestSuite) Test_ShouldReturnCreatedArtifact_WhenNoError() {
	//GIVEN
	var expectedArtifactDto = model.ArtifactDTO{}
	var expectedError error = nil

	//WHEN
	var createdArtifactDto, err = suite.artifactService.CreateArtifact()

	//THEN
	suite.Assert().Equal(expectedError, err)
	suite.Assert().Equal(&expectedArtifactDto, createdArtifactDto)
}

func (suite *ArtifactCreationTestSuite) Test_ShouldReturnEncounteredError_WhenDaoFails() {
	//GIVEN
	var expectedArtifactDto *model.ArtifactDTO = nil
	var expectedError = errors.New("Any Dao Error")

	suite.mockedArtifactDao.expectedError = expectedError

	//WHEN
	var createdArtifactDto, err = suite.artifactService.CreateArtifact()

	//THEN
	suite.Assert().Equal(expectedError, err)
	suite.Assert().Equal(expectedArtifactDto, createdArtifactDto)
}

func (suite *ArtifactCreationTestSuite) Test_ShouldReturnEncounteredError_WhenUUIDGenerationFails() {
	//GIVEN
	var expectedArtifactDto *model.ArtifactDTO = nil
	var expectedError = errors.New("Any UUID generation Error")

	suite.mockedUtilsService.NewUUID_ExpectedError = expectedError

	//WHEN
	var createdArtifactDto, err = suite.artifactService.CreateArtifact()

	//THEN
	suite.Assert().Equal(expectedError, err)
	suite.Assert().Equal(expectedArtifactDto, createdArtifactDto)
}
