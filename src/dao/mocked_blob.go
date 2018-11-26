package dao

import (
	"github.com/UPrefer/StorageService/model"
	"github.com/stretchr/testify/mock"
	"io"
)

type MockedBlobDao struct {
	mock.Mock

	ExpectedSaveDataError error
}

func (dao *MockedBlobDao) SaveData(dto *model.ArtifactDTO, contentType string, reader io.Reader) error {
	return dao.ExpectedSaveDataError
}
