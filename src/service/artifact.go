package service

import (
	"github.com/UPrefer/StorageService/dao"
	"github.com/UPrefer/StorageService/model"
)

type IArtifactService interface {
	CreateArtifact() (*model.ArtifactDTO, error)
	ReadArtifact(id string) (*model.ArtifactDTO, error)
}

func NewArtifactService(utilsService IUtilsService, artifactDAO dao.IArtifactDao) *ArtifactService {
	return &ArtifactService{utilsService: utilsService, artifactDao: artifactDAO}
}

type ArtifactService struct {
	utilsService IUtilsService
	artifactDao  dao.IArtifactDao
}

func (artifactService *ArtifactService) ReadArtifact(id string) (*model.ArtifactDTO, error) {
	panic("implement me")
}

func (artifactService *ArtifactService) CreateArtifact() (*model.ArtifactDTO, error) {
	var newId, err = artifactService.utilsService.NewUUID()
	if err != nil {
		return nil, err
	}

	var artifactDto = &model.ArtifactDTO{Uuid: newId}

	err = artifactService.artifactDao.CreateArtifact(artifactDto)
	if err != nil {
		return nil, err
	}
	return artifactDto, nil
}
