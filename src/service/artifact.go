package service

import (
	"github.com/UPrefer/StorageService/dao"
	"github.com/UPrefer/StorageService/model"
)

type IArtifactService interface {
	CreateArtifact(artifactDto *model.ArtifactDTO) (*model.ArtifactDTO, error)
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
	var (
		err      error
		artifact *model.ArtifactDTO
	)

	artifact, err = artifactService.artifactDao.FindWaitingForUploadArtifact(id)
	if err != nil {
		return nil, err
	}

	if artifact == nil {
		artifact, err = artifactService.artifactDao.FindUploadedArtifact(id)
	}

	return artifact, err
}

func (artifactService *ArtifactService) CreateArtifact(artifactDto *model.ArtifactDTO) (*model.ArtifactDTO, error) {
	var newId, err = artifactService.utilsService.NewUUID()
	if err != nil {
		return nil, err
	}

	var newArtifactDto = &model.ArtifactDTO{Uuid: newId, Name: artifactDto.Name}

	err = artifactService.artifactDao.CreateArtifact(newArtifactDto)
	if err != nil {
		return newArtifactDto, err
	}
	return newArtifactDto, nil
}
