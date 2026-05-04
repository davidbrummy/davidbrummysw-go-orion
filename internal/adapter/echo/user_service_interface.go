package echo

import (
	"context"

	"github.com/davidbrummysw/davidbrummysw-go-orion/internal/core/domain"
)

type UserServiceInterface interface {
	Test(ctx context.Context) (*domain.User, error)
}
