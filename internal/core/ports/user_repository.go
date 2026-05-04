package ports

import (
	"context"

	"github.com/davidbrummysw/davidbrummysw-go-orion/internal/core/domain"
)

type UserRepository interface {
	GetUser(ctx context.Context) (*domain.User, error)
}
