package converter

import (
	"github.com/celtic93/auth/internal/model"
	desc "github.com/celtic93/auth/pkg/v1/user"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToGetResponseFromUser(user *model.User) *desc.GetResponse {
	return &desc.GetResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      desc.Role(user.Role),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}

func ToUserFromCreateRequest(req *desc.CreateRequest) *model.User {
	return &model.User{
		Name:                 req.Name,
		Email:                req.Email,
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
		Role:                 int32(req.GetRole()),
	}
}

func ToUserFromUpdateRequest(req *desc.UpdateRequest) *model.User {
	return &model.User{
		ID:    req.Id,
		Name:  *req.Name,
		Email: req.Email,
	}
}
