package service

import (
	"github.com/UPrefer/StorageService/dao"
	"github.com/UPrefer/StorageService/model"
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
	var err = artifactService.artifactDao.CreateArtifact(artifactDto)
	if err != nil {
		return nil, err
	}
	return artifactDto, nil
}
