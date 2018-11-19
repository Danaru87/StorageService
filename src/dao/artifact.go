package dao

import "github.com/UPrefer/StorageService/model"

type IArtifactDao interface {
	CreateArtifact(*model.ArtifactDTO) error
}

type ArtifactDao struct{}

func NewArtifactDao() *ArtifactDao {
	return &ArtifactDao{}
}

func (*ArtifactDao) CreateArtifact(artifactDto *model.ArtifactDTO) error {
	return nil
}
