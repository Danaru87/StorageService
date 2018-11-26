package service

import (
	"errors"
	"github.com/UPrefer/StorageService/dao"
	"github.com/UPrefer/StorageService/model"
	"github.com/globalsign/mgo"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ArtifactReadTestSuite struct {
	suite.Suite

	artifactId string

	mockedUtilsService *MockedUtilsService
	mockedArtifactDao  *dao.MockedArtifactDao
	artifactService    *ArtifactService
}

func (suite *ArtifactReadTestSuite) SetupTest() {
	suite.artifactId = "artifactId"
	suite.mockedUtilsService = &MockedUtilsService{}
	suite.mockedArtifactDao = &dao.MockedArtifactDao{}
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

	suite.mockedArtifactDao.ExpectedWaitingForUploadArtifact = expectedArtifact

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

	suite.mockedArtifactDao.ExpectedWaitingForUploadError = expectedError

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

	suite.mockedArtifactDao.ExpectedWaitingForUploadError = mgo.ErrNotFound
	suite.mockedArtifactDao.ExpectedAlreadyUploadedArtifact = expectedArtifact

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

	suite.mockedArtifactDao.ExpectedWaitingForUploadError = mgo.ErrNotFound
	suite.mockedArtifactDao.ExpectedAlreadyUploadedError = expectedError

	//WHEN
	var actualArtifact, actualError = suite.artifactService.ReadArtifact(suite.artifactId)

	//THEN
	suite.Equal(expectedArtifact, actualArtifact)
	suite.Equal(expectedError, actualError)
}
