package main

import (
	"context"
	"log"
	"net"

	"github.com/fatih/color"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	userAPI "github.com/celtic93/auth/internal/api/user"
	"github.com/celtic93/auth/internal/config"
	userRepository "github.com/celtic93/auth/internal/repository/user"
	userService "github.com/celtic93/auth/internal/service/user"
	desc "github.com/celtic93/auth/pkg/v1/user"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := config.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	pgConfig, err := config.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	conn, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatal(color.RedString("failed to serve grpc server: %v", err))
	}

	ctx := context.Background()
	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	userRepo := userRepository.NewRepository(pool)
	userServ := userService.NewService(userRepo)

	log.Print(color.GreenString("UserAPI grpc server listening on: %s", conn.Addr().String()))

	gsrv := grpc.NewServer()
	reflection.Register(gsrv)

	desc.RegisterUserV1Server(gsrv, userAPI.NewImplementation(userServ))

	if err = gsrv.Serve(conn); err != nil {
		log.Fatal(color.RedString("failed to serve grpc server: %v", err))
	}
}
