package service

import (
	"crypto/rand"
	"errors"
	"github.com/UPrefer/StorageService/mocks"
	"github.com/UPrefer/StorageService/model"
	"github.com/stretchr/testify/suite"
	"io"
	"io/ioutil"
	"testing"
)

type BlobCreationTestSuite struct {
	suite.Suite

	randomReadCloser io.ReadCloser

	blobDao     *mocks.IBlobDao
	artifactDao *mocks.IArtifactDao

	blobService *BlobService
}

func (suite *BlobCreationTestSuite) SetupTest() {
	suite.randomReadCloser = ioutil.NopCloser(io.LimitReader(rand.Reader, 255))
	suite.blobDao = &mocks.IBlobDao{}
	suite.artifactDao = &mocks.IArtifactDao{}
	suite.blobService = &BlobService{artifactDao: suite.artifactDao, blobDao: suite.blobDao}
}

func TestBlobCreation(t *testing.T) {
	suite.Run(t, new(BlobCreationTestSuite))
}

func (suite *BlobCreationTestSuite) Test_ShouldSaveData_AndDeleteArtifactPlaceholder_WhenArtifactWaitsForUpload_AndNoError() {
	//GIVEN
	var (
		expectedError error = nil
		contentType         = "application/blob"
		artifactId          = "artifact1"
		artifact            = &model.ArtifactDTO{Uuid: artifactId}
	)

	suite.artifactDao.On("FindWaitingForUploadArtifact", artifactId).Return(artifact, nil)
	suite.artifactDao.On("FindUploadedArtifact", artifactId).Return(nil, nil)
	suite.blobDao.On("SaveData", artifact, contentType, suite.randomReadCloser).Return(nil)
	suite.artifactDao.On("DeleteWaitingForUploadArtifact", artifactId).Return(nil)

	//WHEN
	var actualError = suite.blobService.SaveBlob(artifactId, contentType, suite.randomReadCloser)

	//THEN
	suite.Equal(expectedError, actualError)
	suite.blobDao.AssertCalled(suite.T(), "SaveData", artifact, contentType, suite.randomReadCloser)
	suite.artifactDao.AssertCalled(suite.T(), "DeleteWaitingForUploadArtifact", artifactId)
}

func (suite *BlobCreationTestSuite) Test_ShouldReturnArtifactNotFoundError_WhenArtifactNotFound() {
	//GIVEN
	var (
		artifactId    = "artifact1"
		expectedError = ErrArtifactNofFound
	)
	suite.artifactDao.On("FindWaitingForUploadArtifact", artifactId).Return(nil, nil)
	suite.artifactDao.On("FindUploadedArtifact", artifactId).Return(nil, nil)

	//WHEN
	var actualError = suite.blobService.SaveBlob(artifactId, "", suite.randomReadCloser)

	//THEN
	suite.Equal(expectedError, actualError)
}

func (suite *BlobCreationTestSuite) Test_ShouldReturnArtifactAlreadyUploadedError_WhenArtifactAlreadyUploaded() {
	//GIVEN
	var (
		artifactId    = "artifact1"
		expectedError = ErrArtifactAlreadyUploaded
	)
	suite.artifactDao.On("FindWaitingForUploadArtifact", artifactId).Return(nil, nil)
	suite.artifactDao.On("FindUploadedArtifact", artifactId).Return(&model.ArtifactDTO{}, nil)

	//WHEN
	var actualError = suite.blobService.SaveBlob(artifactId, "", suite.randomReadCloser)

	//THEN
	suite.Equal(expectedError, actualError)
}

func (suite *BlobCreationTestSuite) Test_ShouldReturnError_WhenSaveDataFails() {
	//GIVEN
	var (
		expectedError = errors.New("unexpected error")
		artifactId    = "artifactId"
		artifact      = &model.ArtifactDTO{Uuid: artifactId}
		contentType   = "application/data"
	)

	suite.artifactDao.On("FindWaitingForUploadArtifact", artifactId).Return(artifact, nil)
	suite.artifactDao.On("FindUploadedArtifact", artifactId).Return(nil, nil)
	suite.blobDao.On("SaveData", artifact, contentType, suite.randomReadCloser).Return(expectedError)

	//WHEN
	var actualError = suite.blobService.SaveBlob(artifactId, contentType, suite.randomReadCloser)

	//THEN
	suite.Equal(expectedError, actualError)
}

func (suite *BlobCreationTestSuite) Test_ShouldReturnError_WhenDeleteArtifactPlaceholderFails() {
	//GIVEN
	var (
		expectedError = errors.New("delete artifact placeholder error")
		artifactId    = "artifact1"
		artifact      = &model.ArtifactDTO{Uuid: artifactId}
		contentType   = "application/data"
	)

	suite.artifactDao.On("FindWaitingForUploadArtifact", artifactId).Return(artifact, nil)
	suite.artifactDao.On("FindUploadedArtifact", artifactId).Return(nil, nil)
	suite.blobDao.On("SaveData", artifact, contentType, suite.randomReadCloser).Return(nil)
	suite.artifactDao.On("DeleteWaitingForUploadArtifact", artifactId).Return(expectedError)

	//WHEN
	var actualError = suite.blobService.SaveBlob(artifactId, contentType, suite.randomReadCloser)

	//THEN
	suite.Equal(expectedError, actualError)
}
