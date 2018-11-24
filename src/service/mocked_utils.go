package service

type MockedUtilsService struct {
	NewUUID_ExpectedUUID  string
	NewUUID_ExpectedError error
}

func (service *MockedUtilsService) NewUUID() (string, error) {
	return service.NewUUID_ExpectedUUID, service.NewUUID_ExpectedError
}
