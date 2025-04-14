package user

import (
	"context"
	"log"

	"github.com/celtic93/auth/internal/model"
)

func (s *serv) Get(ctx context.Context, id int64) (*model.User, error) {
	log.Printf("service.User.Get started. Get user with id: %d", id)
	user, err := s.userRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	log.Printf("service.User.Get ended. Got user with id: %d", id)

	return user, nil
}
