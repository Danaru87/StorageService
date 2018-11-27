package service

import (
	"github.com/UPrefer/StorageService/dao"
	"io"
	"log"
)

type IBlobService interface {
	SaveBlob(artifactID string, contentType string, reader io.ReadCloser) error
	ReadBlob(artifactId string) (contentType string, length int64, reader io.ReadCloser, err error)
}

func NewBlobService(blobDao dao.IBlobDao, artifactDao dao.IArtifactDao) *BlobService {
	return &BlobService{blobDao: blobDao, artifactDao: artifactDao}
}

type BlobService struct {
	blobDao     dao.IBlobDao
	artifactDao dao.IArtifactDao
}

func (blobService *BlobService) ReadBlob(artifactId string) (contentType string, length int64, reader io.ReadCloser, resultError error) {

	artifactDto, resultError := blobService.artifactDao.FindUploadedArtifact(artifactId)

	if artifactDto == nil {
		resultError = ErrArtifactNotFound
	}

	if resultError == nil {
		reader, resultError = blobService.blobDao.ReadData(artifactId)

		if reader == nil {
			resultError = ErrArtifactNotFound
		}
	}

	if resultError != nil {
		return "", 0, nil, resultError
	}

	return artifactDto.ContentType, artifactDto.Length, reader, resultError
}

func (blobService *BlobService) SaveBlob(artifactId string, contentType string, reader io.ReadCloser) error {
	defer reader.Close()
	var err error

	alreadyUploadedArtifact, err := blobService.artifactDao.FindUploadedArtifact(artifactId)
	if alreadyUploadedArtifact != nil {
		log.Print(err)
		return ErrArtifactAlreadyUploaded
	}

	waitingForUploadArtifact, err := blobService.artifactDao.FindWaitingForUploadArtifact(artifactId)
	if waitingForUploadArtifact == nil {
		log.Print(err)
		return ErrArtifactNotFound
	}

	err = blobService.blobDao.SaveData(waitingForUploadArtifact, contentType, reader)
	if err != nil {
		return err
	}

	return blobService.artifactDao.DeleteWaitingForUploadArtifact(artifactId)
}
