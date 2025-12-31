package echo

import "github.com/davidbrummysw/davidbrummysw-go-orion/internal/core/domain"

type UserServiceInterface interface {
	Test() *domain.User
}
