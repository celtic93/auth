package user

import (
	"github.com/celtic93/auth/internal/repository"
	"github.com/celtic93/auth/internal/service"
)

type serv struct {
	userRepository repository.UserRepository
}

func NewService(userRepository repository.UserRepository) service.UserService {
	return &serv{
		userRepository: userRepository,
	}
}
