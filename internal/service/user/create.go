package user

import (
	"context"

	"github.com/celtic93/auth/internal/model"
)

func (s *serv) Create(ctx context.Context, user *model.User) (int64, error) {
	id, err := s.userRepository.Create(ctx, user)
	if err != nil {
		return 0, err
	}

	return id, nil
}
