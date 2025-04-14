package user

import (
	"context"
	"log"

	"github.com/celtic93/auth/internal/api/user/converter"
	desc "github.com/celtic93/auth/pkg/v1/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("api.User.Update started. Update user with id: %d", req.Id)
	err := i.userService.Update(ctx, converter.ToUserFromUpdateRequest(req))
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Printf("api.User.Update ended. Updated user with id: %d", req.Id)

	return &emptypb.Empty{}, nil
}
