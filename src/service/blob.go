package service

import (
	"github.com/UPrefer/StorageService/dao"
	"io"
	"log"
)

type IBlobService interface {
	SaveBlob(artifactID string, contentType string, reader io.ReadCloser) error
}

func NewBlobService(blobDao dao.IBlobDao) *BlobService {
	return &BlobService{blobDao: blobDao}
}

type BlobService struct {
	blobDao     dao.IBlobDao
	artifactDao dao.IArtifactDao
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
		return ErrArtifactNofFound
	}

	err = blobService.blobDao.SaveData(waitingForUploadArtifact, contentType, reader)
	if err != nil {
		return err
	}

	return blobService.artifactDao.DeleteWaitingForUploadArtifact(artifactId)
}
