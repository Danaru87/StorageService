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

type BlobReadTestSuite struct {
	suite.Suite

	reader io.ReadCloser

	blobDao     *mocks.IBlobDao
	artifactDao *mocks.IArtifactDao

	blobService *BlobService
}

func (suite *BlobReadTestSuite) SetupTest() {
	suite.reader = ioutil.NopCloser(io.LimitReader(rand.Reader, 255))

	suite.artifactDao = &mocks.IArtifactDao{}
	suite.blobDao = &mocks.IBlobDao{}

	suite.blobService = &BlobService{artifactDao: suite.artifactDao, blobDao: suite.blobDao}
}

func TestBlobRead(t *testing.T) {
	suite.Run(t, new(BlobReadTestSuite))
}

func (suite *BlobReadTestSuite) Test_ShouldReturnFile_AndContentType_AndLength_WhenFindSucceeds() {
	//GIVEN
	var (
		expectedContentType       = "application/data"
		expectedLength            = int64(12)
		expectedError       error = nil
		artifactId                = "artifact1"
	)

	suite.artifactDao.On("FindUploadedArtifact", artifactId).Return(&model.ArtifactDTO{Uuid: artifactId, ContentType: expectedContentType, Length: expectedLength}, nil)
	suite.blobDao.On("ReadData", artifactId).Return(suite.reader, nil)

	//WHEN
	var actualContentType, actualLength, actualReader, actualError = suite.blobService.ReadBlob(artifactId)

	//THEN
	suite.Equal(suite.reader, actualReader)
	suite.Equal(expectedLength, actualLength)
	suite.Equal(expectedContentType, actualContentType)
	suite.Equal(expectedError, actualError)
}

func (suite *BlobReadTestSuite) Test_ShouldReturnArtifactNotFound_WhenArtifactNotFound() {
	//GIVEN
	var (
		expectedError = ErrArtifactNotFound
		artifactId    = "artifact1"
	)

	suite.artifactDao.On("FindUploadedArtifact", artifactId).Return(nil, nil)
	suite.blobDao.On("ReadData", artifactId).Return(suite.reader, nil)

	//WHEN
	var _, _, _, actualError = suite.blobService.ReadBlob(artifactId)

	//THEN
	suite.Equal(expectedError, actualError)
}

func (suite *BlobReadTestSuite) Test_ShouldReturnArtifactNotFound_WhenBlobNotUploaded() {
	//GIVEN
	var (
		expectedError = ErrArtifactNotFound
		artifactId    = "artifact1"
	)

	suite.artifactDao.On("FindUploadedArtifact", artifactId).Return(&model.ArtifactDTO{Length: 0, ContentType: "", Uuid: artifactId}, nil)
	suite.blobDao.On("ReadData", artifactId).Return(nil, nil)

	//WHEN
	var _, _, _, actualError = suite.blobService.ReadBlob(artifactId)

	//THEN
	suite.Equal(expectedError, actualError)
}

func (suite *BlobReadTestSuite) Test_ShouldReturnError_WhenFindUploadedArtifactFails() {
	//GIVEN
	var (
		expectedError = errors.New("random error")
		artifactId    = "artifact1"
	)

	suite.artifactDao.On("FindUploadedArtifact", artifactId).Return(&model.ArtifactDTO{Length: 0, ContentType: "", Uuid: artifactId}, expectedError)
	suite.blobDao.On("ReadData", artifactId).Return(nil, nil)

	//WHEN
	var _, _, _, actualError = suite.blobService.ReadBlob(artifactId)

	//THEN
	suite.Equal(expectedError, actualError)
}

func (suite *BlobReadTestSuite) Test_ShouldReturnError_WhenReadBlobFails() {
	//GIVEN
	var (
		expectedError = errors.New("random error")
		artifactId    = "artifact1"
	)

	suite.artifactDao.On("FindUploadedArtifact", artifactId).Return(&model.ArtifactDTO{Length: 0, ContentType: "", Uuid: artifactId}, nil)
	suite.blobDao.On("ReadData", artifactId).Return(suite.reader, expectedError)

	//WHEN
	var _, _, _, actualError = suite.blobService.ReadBlob(artifactId)

	//THEN
	suite.Equal(expectedError, actualError)
}
