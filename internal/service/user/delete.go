package user

import (
	"context"
	"log"
)

func (s *serv) Delete(ctx context.Context, id int64) error {
	log.Printf("service.User.Delete started. Delete user with id: %d", id)
	err := s.userRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	log.Printf("service.User.Delete ended. Deleted user with id: %d", id)

	return nil
}
