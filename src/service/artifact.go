package service

import (
	"github.com/UPrefer/StorageService/model"
)

type IArtifactService interface {
	CreateArtifact(artifactDto *model.ArtifactDTO) (*model.ArtifactDTO, error)
}

type ArtifactService struct{}

func (ArtifactService) CreateArtifact(artifactDto *model.ArtifactDTO) (*model.ArtifactDTO, error) {
	panic("implement me")
}
