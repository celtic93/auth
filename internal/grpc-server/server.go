package server

import (
	"context"
	"log"

	"github.com/brianvoe/gofakeit"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/celtic93/auth/pkg/v1/user"
)

type Server struct {
	desc.UnimplementedUserV1Server
	Pool *pgxpool.Pool
}

// Get: gets user by ID
func (s *Server) Get(_ context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("server.Get User id: %d", req.GetId())

	return &desc.GetResponse{
		Id:        req.GetId(),
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Role:      desc.Role_USER,
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}, nil
}

// Create: creates user
func (s *Server) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("server.Create User id: %s", req.GetEmail())

	return &desc.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

// Update: updates user
func (s *Server) Update(_ context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("server.Update User id: %d", req.GetId())

	return &emptypb.Empty{}, nil
}

// Delete: deletes user
func (s *Server) Delete(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("server.Delete User id: %d", req.GetId())

	return &emptypb.Empty{}, nil
}
