package service

import (
	"crypto/rand"
	"github.com/UPrefer/StorageService/dao"
	"github.com/UPrefer/StorageService/model"
	"github.com/stretchr/testify/suite"
	"io"
	"io/ioutil"
	"testing"
)

type BlobCreationTestSuite struct {
	suite.Suite

	randomReadCloser io.ReadCloser

	blobDao     *dao.MockedBlobDao
	artifactDao *dao.MockedArtifactDao

	blobService *BlobService
}

func (suite *BlobCreationTestSuite) SetupTest() {
	suite.randomReadCloser = ioutil.NopCloser(io.LimitReader(rand.Reader, 255))
	suite.blobDao = &dao.MockedBlobDao{}
	suite.artifactDao = &dao.MockedArtifactDao{}
	suite.blobService = &BlobService{artifactDao: suite.artifactDao, blobDao: suite.blobDao}
}

func TestBlobCreation(t *testing.T) {
	suite.Run(t, new(BlobCreationTestSuite))
}

func (suite *BlobCreationTestSuite) Test_ShouldReturnArtifactNotFoundError_WhenArtifactNotFound() {
	//GIVEN
	var expectedError = ErrArtifactNofFound
	suite.artifactDao.ExpectedWaitingForUploadArtifact = nil

	//WHEN
	var actualError = suite.blobService.SaveBlob("", "", suite.randomReadCloser)

	//THEN
	suite.Equal(expectedError, actualError)
}

func (suite *BlobCreationTestSuite) Test_ShouldReturnArtifactAlreadyUploadedError_WhenArtifactAlreadyUploaded() {
	//GIVEN
	var expectedError = ErrArtifactAlreadyUploaded
	suite.artifactDao.ExpectedWaitingForUploadArtifact = nil
	suite.artifactDao.ExpectedAlreadyUploadedArtifact = &model.ArtifactDTO{}

	//WHEN
	var actualError = suite.blobService.SaveBlob("", "", suite.randomReadCloser)

	//THEN
	suite.Equal(expectedError, actualError)
}

//func (suite *BlobCreationTestSuite) Test_ShouldSaveData_WhenArtifactWaitsForUpload() {
//	//GIVEN
//	var (
//		expectedError error = nil
//		expectedArtifactId  = "adfg"
//		expectedContentType = "application/blob"
//	)
//
//	//WHEN
//	var actualError = suite.blobService.SaveBlob(expectedArtifactId, expectedContentType, suite.randomReadCloser)
//
//	//THEN
//	suite.Equal(expectedError, actualError)
//	suite.blobDao.MethodCalled("SaveData", expectedArtifactId, expectedContentType, suite.randomReadCloser)
//}
