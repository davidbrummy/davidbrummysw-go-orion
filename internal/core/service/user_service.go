package service

import (
	"context"

	"github.com/davidbrummysw/davidbrummysw-go-orion/internal/core/domain"
	"github.com/davidbrummysw/davidbrummysw-go-orion/internal/core/ports"
)

// Adapter implements the DbPort interface
type UserService struct {
	userRepository ports.UserRepository
}

func NewUserService(userRepository ports.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

// Test implements ports.UserServicePort.
func (service *UserService) Test(ctx context.Context) (*domain.User, error) {
	return service.userRepository.GetUser(ctx)
}
