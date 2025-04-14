package user

import (
	"context"
	"log"

	desc "github.com/celtic93/auth/pkg/v1/user"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := i.userService.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	log.Printf("deleted user with id: %d", req.Id)

	return &emptypb.Empty{}, nil
}
