package user

import (
	"github.com/celtic93/auth/internal/repository"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	usersTable string = "users"

	idColumn        string = "id"
	nameColumn      string = "name"
	emailColumn     string = "email"
	passwordColumn  string = "password"
	roleColumn      string = "role"
	createdAtColumn string = "created_at"
	updatedAtColumn string = "updated_at"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.UserRepository {
	return &repo{db: db}
}
