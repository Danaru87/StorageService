package dao

import (
	"github.com/UPrefer/StorageService/config"
	"github.com/UPrefer/StorageService/model"
	"github.com/globalsign/mgo"
	"io"
)

type IBlobDao interface {
	SaveData(dto *model.ArtifactDTO, contentType string, reader io.Reader) error
	ReadData(artifactId string) (io.ReadCloser, error)
}

func NewMongoBlobDao(database *config.Database) *MongoBlobDao {
	return &MongoBlobDao{collectionName: "artifact", database: database}
}

type MongoBlobDao struct {
	collectionName string
	database       *config.Database
}

func (artifactDao *MongoBlobDao) ReadData(artifactId string) (reader io.ReadCloser, err error) {
	reader, err = artifactDao.database.OpenGridFsReader(artifactDao.collectionName, artifactId)

	if err == mgo.ErrNotFound {
		err = nil
		reader = nil
	}
	return reader, err
}

func (artifactDao *MongoBlobDao) SaveData(artifactDto *model.ArtifactDTO, contentType string, reader io.Reader) error {
	var err error = nil

	artifactDao.database.HandleRequest(func(database *mgo.Database) {
		var collection = database.GridFS(artifactDao.collectionName)
		var createdFile, err = collection.Create(artifactDto.Uuid)
		if err == nil {
			createdFile.SetId(artifactDto.Uuid)
			createdFile.SetContentType(contentType)
			createdFile.SetName(artifactDto.Name)

			var _, copyErr = io.Copy(createdFile, reader)

			if copyErr != io.EOF {
				err = copyErr
			}

			err = createdFile.Close()
		}
	})

	return err
}
