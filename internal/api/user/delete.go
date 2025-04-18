package user

import (
	"context"
	"log"

	desc "github.com/celtic93/auth/pkg/v1/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("api.User.Delete started. Delete user with id: %d", req.Id)
	err := i.userService.Delete(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Printf("api.User.Delete ended. Deleted user with id: %d", req.Id)

	return &emptypb.Empty{}, nil
}
