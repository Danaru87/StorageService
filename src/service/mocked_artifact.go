package service

import (
	"github.com/UPrefer/StorageService/model"
	"github.com/stretchr/testify/mock"
)

type MockedArtifactService struct {
	mock.Mock

	CreateArtifact_ExpectedArtifact *model.ArtifactDTO
	CreateArtifact_ExpectedError    error
	ReadArtifact_ExpectedArtifact   *model.ArtifactDTO
	ReadArtifact_ExpectedError      error
}

func (service *MockedArtifactService) ReadArtifact(id string) (*model.ArtifactDTO, error) {
	return service.ReadArtifact_ExpectedArtifact, service.ReadArtifact_ExpectedError
}

func (artifactService *MockedArtifactService) CreateArtifact() (*model.ArtifactDTO, error) {
	return artifactService.CreateArtifact_ExpectedArtifact, artifactService.CreateArtifact_ExpectedError
}
