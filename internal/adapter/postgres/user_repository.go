package postgres

import (
	"context"

	"github.com/davidbrummysw/davidbrummysw-go-orion/internal/core/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (repository *UserRepository) GetUser(ctx context.Context) (*domain.User, error) {
	row := repository.db.QueryRow(ctx, `
		select id, uuid, user_name
		from users
		order by id
		limit 1
	`)

	var id int64
	var uuid string
	var userName string

	if err := row.Scan(&id, &uuid, &userName); err != nil {
		return nil, err
	}

	userID := uint64(id)

	return domain.NewUserFill(&userID, &uuid, &userName), nil
}
