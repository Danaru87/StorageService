package service

import (
	"errors"
	"github.com/UPrefer/StorageService/mocks"
	"github.com/UPrefer/StorageService/model"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ArtifactReadTestSuite struct {
	suite.Suite

	artifactId string

	mockedUtilsService *mocks.IUtilsService
	mockedArtifactDao  *mocks.IArtifactDao
	artifactService    *ArtifactService
}

func (suite *ArtifactReadTestSuite) SetupTest() {
	suite.artifactId = "artifactId"
	suite.mockedUtilsService = &mocks.IUtilsService{}
	suite.mockedArtifactDao = &mocks.IArtifactDao{}
	suite.artifactService = &ArtifactService{utilsService: suite.mockedUtilsService, artifactDao: suite.mockedArtifactDao}
}

func TestArtifactRead(t *testing.T) {
	suite.Run(t, new(ArtifactReadTestSuite))
}

func (suite *ArtifactReadTestSuite) Test_ShouldReturnArtifact_WhenWaitingForUploadSucceeds() {
	//GIVEN
	var (
		expectedArtifact       = &model.ArtifactDTO{Uuid: suite.artifactId}
		expectedError    error = nil
	)

	suite.mockedArtifactDao.On("FindWaitingForUploadArtifact", suite.artifactId).Return(expectedArtifact, expectedError)

	//WHEN
	var actualArtifact, actualError = suite.artifactService.ReadArtifact(suite.artifactId)

	//THEN
	suite.Equal(expectedError, actualError)
	suite.Equal(expectedArtifact, actualArtifact)
}

func (suite *ArtifactReadTestSuite) Test_ShouldReturnError_WhenFindWaitingForUploadFails() {
	//GIVEN
	var (
		expectedArtifact *model.ArtifactDTO = nil
		expectedError                       = errors.New("findWaitingForUploadError")
	)

	suite.mockedArtifactDao.On("FindWaitingForUploadArtifact", suite.artifactId).Return(expectedArtifact, expectedError)

	//WHEN
	var actualArtifact, actualError = suite.artifactService.ReadArtifact(suite.artifactId)

	//THEN
	suite.Equal(expectedError, actualError)
	suite.Equal(expectedArtifact, actualArtifact)
}

func (suite *ArtifactReadTestSuite) Test_ShouldReturnArtifact_WhenFindAlreadyUploadedSucceeds() {
	//GIVEN
	var (
		expectedArtifact       = &model.ArtifactDTO{Uuid: suite.artifactId}
		expectedError    error = nil
	)

	suite.mockedArtifactDao.On("FindWaitingForUploadArtifact", suite.artifactId).Return(nil, nil)
	suite.mockedArtifactDao.On("FindUploadedArtifact", suite.artifactId).Return(expectedArtifact, expectedError)

	//WHEN
	var actualArtifact, actualError = suite.artifactService.ReadArtifact(suite.artifactId)

	//THEN
	suite.Equal(expectedError, actualError)
	suite.Equal(expectedArtifact, actualArtifact)
}

func (suite *ArtifactReadTestSuite) Test_ShouldReturnError_WhenFindAlreadyUploadedFails() {
	//GIVEN
	var (
		expectedArtifact *model.ArtifactDTO = nil
		expectedError                       = errors.New("findAlreadyUploadedError")
	)

	suite.mockedArtifactDao.On("FindWaitingForUploadArtifact", suite.artifactId).Return(nil, nil)
	suite.mockedArtifactDao.On("FindUploadedArtifact", suite.artifactId).Return(expectedArtifact, expectedError)

	//WHEN
	var actualArtifact, actualError = suite.artifactService.ReadArtifact(suite.artifactId)

	//THEN
	suite.Equal(expectedArtifact, actualArtifact)
	suite.Equal(expectedError, actualError)
}

func (suite *ArtifactReadTestSuite) Test_ShouldReturnNil_WhenNoArtifactFound() {
	//GIVEN
	var (
		expectedArtifact *model.ArtifactDTO = nil
		expectedError    error              = nil
	)

	suite.mockedArtifactDao.On("FindWaitingForUploadArtifact", suite.artifactId).Return(nil, nil)
	suite.mockedArtifactDao.On("FindUploadedArtifact", suite.artifactId).Return(nil, nil)

	//WHEN
	var actualArtifact, actualError = suite.artifactService.ReadArtifact(suite.artifactId)

	//THEN
	suite.Equal(expectedArtifact, actualArtifact)
	suite.Equal(expectedError, actualError)
}
