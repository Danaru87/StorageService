package service

import (
	"github.com/UPrefer/StorageService/dao"
	"io"
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

	var alreadyUploadedArtifact, _ = blobService.artifactDao.FindUploadedArtifact(artifactId)
	if alreadyUploadedArtifact != nil {
		return ErrArtifactAlreadyUploaded
	}

	var waitingForUploadArtifact, _ = blobService.artifactDao.FindWaitingForUploadArtifact(artifactId)
	if waitingForUploadArtifact == nil {
		return ErrArtifactNofFound
	}

	return nil
}
