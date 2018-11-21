package service

import (
	"github.com/UPrefer/StorageService/dao"
	"github.com/UPrefer/StorageService/model"
	"github.com/google/uuid"
)

type IArtifactService interface {
	CreateArtifact(artifactDto *model.ArtifactDTO) (*model.ArtifactDTO, error)
}

type ArtifactService struct {
	artifactDao dao.IArtifactDao
}

func NewArtifactService(artifactDAO dao.IArtifactDao) *ArtifactService {
	return &ArtifactService{artifactDao: artifactDAO}
}

func (artifactService *ArtifactService) CreateArtifact(artifactDto *model.ArtifactDTO) (*model.ArtifactDTO, error) {
	var newId, err = uuid.NewUUID()
	if err != nil {
		return artifactDto, err
	}
	artifactDto.Uuid = newId.String()

	err = artifactService.artifactDao.CreateArtifact(artifactDto)
	if err != nil {
		return artifactDto, err
	}
	return artifactDto, nil
}
