package service

import "github.com/google/uuid"

type IUtilsService interface {
	NewUUID() (string, error)
}

func NewUtilsService() *UtilsService {
	return &UtilsService{}
}

type UtilsService struct{}

func (UtilsService) NewUUID() (string, error) {
	var newUuid, err = uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return newUuid.String(), err
}
