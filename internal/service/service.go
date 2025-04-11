package service

import (
	"context"

	"github.com/celtic93/auth/internal/model"
)

type UserService interface {
	Get(ctx context.Context, id int64) (*model.User, error)
	Create(ctx context.Context, user *model.User) (int64, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id int64) error
}
