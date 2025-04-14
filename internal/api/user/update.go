package user

import (
	"context"
	"log"

	"github.com/celtic93/auth/internal/api/user/converter"
	desc "github.com/celtic93/auth/pkg/v1/user"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	err := i.userService.Update(ctx, converter.ToUserFromUpdateRequest(req))
	if err != nil {
		return nil, err
	}

	log.Printf("updated user with id: %d", req.Id)

	return &emptypb.Empty{}, nil
}
