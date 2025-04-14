package user

import (
	"context"
	"log"

	"github.com/celtic93/auth/internal/api/user/converter"
	desc "github.com/celtic93/auth/pkg/v1/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	user, err := i.userService.Get(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Printf("get user with id: %d", req.Id)

	return converter.ToGetResponseFromUser(user), nil
}
