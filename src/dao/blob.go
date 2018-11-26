package dao

import (
	"github.com/UPrefer/StorageService/config"
	"github.com/UPrefer/StorageService/model"
	"github.com/globalsign/mgo"
	"io"
)

type IBlobDao interface {
	SaveData(dto *model.ArtifactDTO, contentType string, reader io.Reader) error
}

func NewMongoBlobDao(database *config.Database) *MongoBlobDao {
	return &MongoBlobDao{collectionName: "artifact", database: database}
}

type MongoBlobDao struct {
	collectionName string
	database       *config.Database
}

func (artifactDao *MongoBlobDao) SaveData(artifactDto *model.ArtifactDTO, contentType string, reader io.Reader) error {
	var err error = nil

	artifactDao.database.HandleRequest(func(database *mgo.Database) {
		var collection = database.GridFS(artifactDao.collectionName)
		var createdFile, err = collection.Create(artifactDto.Uuid)
		if err == nil {
			createdFile.SetId(artifactDto.Uuid)
			createdFile.SetContentType(contentType)

			var fileChunk = make([]byte, 255)
			var readErr error
			for readErr == nil {
				_, readErr = reader.Read(fileChunk)
				createdFile.Write(fileChunk)
			}

			if readErr != io.EOF {
				err = readErr
			}

			err = createdFile.Close()
		}
	})

	return err
}
