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

func NewArtifactDao(database *config.Database) *ArtifactDao {
	return &ArtifactDao{uploadedCollectionName: "artifact", waitingForUploadCollectionName: "artifact.waitingForUpload", database: database}
}

type ArtifactDao struct {
	waitingForUploadCollectionName string
	uploadedCollectionName         string
	database                       *config.Database
}

func (artifactDao *ArtifactDao) FindUploadedArtifact(id string) (*model.ArtifactDTO, error) {
	var (
		foundArtifact = &model.ArtifactDTO{}
		err           error
	)
	artifactDao.database.HandleRequest(func(database *mgo.Database) {
		err = database.C(artifactDao.uploadedCollectionName).FindId(id).One(foundArtifact)
	})
	return foundArtifact, err
}

func (artifactDao *ArtifactDao) FindWaitingForUploadArtifact(id string) (*model.ArtifactDTO, error) {
	var (
		foundArtifact = &model.ArtifactDTO{}
		err           error
	)
	artifactDao.database.HandleRequest(func(database *mgo.Database) {
		err = database.C(artifactDao.waitingForUploadCollectionName).FindId(id).One(foundArtifact)
	})
	return foundArtifact, err
}

func (artifactDao *ArtifactDao) CreateArtifact(artifactDto *model.ArtifactDTO) error {
	var err error = nil
	artifactDao.database.HandleRequest(func(database *mgo.Database) {
		err = database.C(artifactDao.waitingForUploadCollectionName).Insert(artifactDto)
	})
	return err
}
