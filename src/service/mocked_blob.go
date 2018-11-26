package service

import (
	"github.com/stretchr/testify/mock"
	"io"
)

type MockedBlobService struct {
	mock.Mock

	ExpectedSaveBlobError error
}

func (mock *MockedBlobService) SaveBlob(artifactId string, contentType string, reader io.ReadCloser) error {
	return mock.ExpectedSaveBlobError
}
