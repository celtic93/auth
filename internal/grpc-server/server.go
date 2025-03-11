package server

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	sq "github.com/Masterminds/squirrel"
	desc "github.com/celtic93/auth/pkg/v1/user"
)

const (
	columnID   string = "id"
	name       string = "name"
	email      string = "email"
	password   string = "password"
	role       string = "role"
	createdAt  string = "created_at"
	updatedAt  string = "updated_at"
	usersTable string = "users"
)

type Server struct {
	desc.UnimplementedUserV1Server
	Pool *pgxpool.Pool
}

// Get: gets user by ID
func (s *Server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("server.Get User id: %d", req.GetId())
	builderSelectOne := sq.Select(columnID, name, email, role, createdAt, updatedAt).
		From(usersTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.GetId()}).
		Limit(1)

	query, args, err := builderSelectOne.ToSql()
	if err != nil {
		log.Print(err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var resp desc.GetResponse
	var createdAt, updatedAt time.Time

	err = s.Pool.QueryRow(ctx, query, args...).Scan(&resp.Id, &resp.Name, &resp.Email, &resp.Role, &createdAt, &updatedAt)
	if err != nil {
		log.Print(err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp.CreatedAt = timestamppb.New(createdAt)
	resp.UpdatedAt = timestamppb.New(updatedAt)
	log.Printf("get user id: %d, name: %s, email: %s, role: %d, created_at: %v, updated_at: %v\n",
		resp.Id, resp.Name, resp.Email, resp.Role, resp.CreatedAt, resp.UpdatedAt)

	return &resp, nil
}

// Create: creates user
func (s *Server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("server.Create User id: %s", req.GetEmail())

	if req.GetPassword() != req.GetPasswordConfirmation() {
		log.Print("password doesn't match")
		return nil, status.Error(codes.InvalidArgument, "password doesn't match")
	}

	timeNow := time.Now()
	builderInsert := sq.Insert(usersTable).
		PlaceholderFormat(sq.Dollar).
		Columns(name, email, password, role, createdAt, updatedAt).
		Values(req.GetName(), req.GetEmail(), req.GetPassword(), req.GetRole(), timeNow, timeNow).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Print(err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var userID int64
	err = s.Pool.QueryRow(ctx, query, args...).Scan(&userID)
	if err != nil {
		log.Print(err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Printf("inserted user with id: %d", userID)

	return &desc.CreateResponse{
		Id: userID,
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
