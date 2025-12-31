package service

import "github.com/davidbrummysw/davidbrummysw-go-orion/internal/core/domain"

// Adapter implements the DbPort interface
type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

// Test implements ports.UserServicePort.
func (service *UserService) Test() *domain.User {

	name := "Test Name"
	uuid := "Test UUID"
	id := uint64(1)

	return domain.NewUserFill(&id, &uuid, &name)
}
