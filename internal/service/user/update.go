package user

import (
	"context"
	"fmt"
	"log"

	"github.com/celtic93/auth/internal/model"
)

func (s *serv) Update(ctx context.Context, user *model.User) error {
	if user.Email == "" {
		log.Print("email is empty")
		return fmt.Errorf("email can't be blank")
	}

	err := s.userRepository.Update(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
