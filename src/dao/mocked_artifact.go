package dao

import "github.com/UPrefer/StorageService/model"

type MockedArtifactDao struct {
	ExpectedError                    error
	ExpectedWaitingForUploadArtifact *model.ArtifactDTO
	ExpectedWaitingForUploadError    error
	ExpectedAlreadyUploadedArtifact  *model.ArtifactDTO
	ExpectedAlreadyUploadedError     error
}

func (dao *MockedArtifactDao) FindUploadedArtifact(id string) (*model.ArtifactDTO, error) {
	return dao.ExpectedAlreadyUploadedArtifact, dao.ExpectedAlreadyUploadedError
}

func (dao *MockedArtifactDao) CreateArtifact(*model.ArtifactDTO) error {
	return dao.ExpectedError
}

func (dao *MockedArtifactDao) FindWaitingForUploadArtifact(s string) (*model.ArtifactDTO, error) {
	return dao.ExpectedWaitingForUploadArtifact, dao.ExpectedWaitingForUploadError
}
