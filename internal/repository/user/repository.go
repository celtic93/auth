package user

import (
	"context"

	"github.com/celtic93/auth/internal/model"
	"github.com/celtic93/auth/internal/repository"
	"github.com/jackc/pgx/v4/pgxpool"
)

// const (
// 	usersTable      string = "users"

// 	idColumn        string = "id"
// 	nameColumn      string = "name"
// 	emailColumn     string = "email"
// 	passwordColumn  string = "password"
// 	roleColumn      string = "role"
// 	createdAtColumn string = "created_at"
// 	updatedAtColumn string = "updated_at"
// )

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.UserRepository {
	return &repo{db: db}
}

func (r *repo) Get(_ context.Context, _ int64) (*model.User, error) {
	return &model.User{}, nil
}

func (r *repo) Create(_ context.Context, _ *model.User) (int64, error) {
	return 0, nil
}

func (r *repo) Update(_ context.Context, _ *model.User) error {
	return nil
}

func (r *repo) Delete(_ context.Context, _ int64) error {
	return nil
}
