package user

import (
	"context"
	"log"

	"github.com/celtic93/auth/internal/api/user/converter"
	desc "github.com/celtic93/auth/pkg/v1/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("api.User.Create started. Create user with id: %s", req.Email)
	id, err := i.userService.Create(ctx, converter.ToUserFromCreateRequest(req))
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Printf("api.User.Create ended. Created user with id: %d", id)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
