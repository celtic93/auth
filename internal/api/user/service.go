package user

import (
	"github.com/celtic93/auth/internal/service"
	desc "github.com/celtic93/auth/pkg/v1/user"
)

type Implementation struct {
	desc.UnimplementedUserV1Server
	userService service.UserService
}

func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
