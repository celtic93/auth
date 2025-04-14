package user

import (
	"context"
	"fmt"
	"log"

	"github.com/celtic93/auth/internal/model"
)

func (s *serv) Create(ctx context.Context, user *model.User) (int64, error) {
	log.Printf("service.User.Create started. Create user with id: %d", user.ID)
	if user.Password != user.PasswordConfirmation {
		log.Print("Error. Password doesn't match")
		return 0, fmt.Errorf("password doesn't match")
	}

	id, err := s.userRepository.Create(ctx, user)
	if err != nil {
		return 0, err
	}

	log.Printf("service.User.Create ended. Created user with id: %d", user.ID)

	return id, nil
}
