package dao

import (
	"github.com/UPrefer/StorageService/config"
	"github.com/UPrefer/StorageService/model"
	"github.com/globalsign/mgo"
)

type IArtifactDao interface {
	CreateArtifact(*model.ArtifactDTO) error
	FindWaitingForUploadArtifact(s string) (*model.ArtifactDTO, error)
	FindUploadedArtifact(id string) (*model.ArtifactDTO, error)
}

type ArtifactDao struct {
	collectionName string
	database       *config.Database
}

func (artifactDao *ArtifactDao) FindUploadedArtifact(id string) (*model.ArtifactDTO, error) {
	panic("implement me")
}

func (artifactDao *ArtifactDao) FindWaitingForUploadArtifact(s string) (*model.ArtifactDTO, error) {
	panic("implement me")
}

func NewArtifactDao(database *config.Database) *ArtifactDao {
	return &ArtifactDao{collectionName: "artifact", database: database}
}

func (artifactDao *ArtifactDao) CreateArtifact(artifactDto *model.ArtifactDTO) error {
	var err error = nil
	artifactDao.database.HandleRequest(func(database *mgo.Database) {
		err = database.C(artifactDao.collectionName).Insert(artifactDto)
	})
	return err
}
