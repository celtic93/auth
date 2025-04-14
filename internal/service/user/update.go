package user

import (
	"context"
	"fmt"
	"log"

	"github.com/celtic93/auth/internal/model"
)

func (s *serv) Update(ctx context.Context, user *model.User) error {
	log.Printf("service.User.Update started. Update user with id: %d", user.ID)
	if user.Email == "" {
		log.Print("Error. Email is empty")
		return fmt.Errorf("email can't be blank")
	}

	err := s.userRepository.Update(ctx, user)
	if err != nil {
		return err
	}

	log.Printf("service.User.Update ended. Updated user with id: %d", user.ID)

	return nil
}
