package app

import (
	"context"
	"log"

	userAPI "github.com/celtic93/auth/internal/api/user"
	"github.com/celtic93/auth/internal/closer"
	"github.com/celtic93/auth/internal/config"
	"github.com/celtic93/auth/internal/repository"
	userRepository "github.com/celtic93/auth/internal/repository/user"
	"github.com/celtic93/auth/internal/service"
	userService "github.com/celtic93/auth/internal/service/user"
	"github.com/jackc/pgx/v4/pgxpool"
)

type serviceProvider struct {
	userImplementation *userAPI.Implementation
	userRepository     repository.UserRepository
	userService        service.UserService

	grpcConfig config.GRPCConfig

	pgConfig config.PGConfig
	db       *pgxpool.Pool
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (sp *serviceProvider) DBClient(ctx context.Context) *pgxpool.Pool {
	if sp.db == nil {
		pool, err := pgxpool.Connect(ctx, sp.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		err = pool.Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %v", err)
		}
		closer.Add(func() error {
			pool.Close()

			return nil
		})

		sp.db = pool
	}

	return sp.db
}

func (sp *serviceProvider) GRPCConfig() config.GRPCConfig {
	if sp.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %v", err)
		}

		sp.grpcConfig = cfg
	}

	return sp.grpcConfig
}

func (sp *serviceProvider) PGConfig() config.PGConfig {
	if sp.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %v", err)
		}

		sp.pgConfig = cfg
	}

	return sp.pgConfig
}

func (sp *serviceProvider) UserImplementation(ctx context.Context) *userAPI.Implementation {
	if sp.userImplementation == nil {
		sp.userImplementation = userAPI.NewImplementation(sp.UserService(ctx))
	}

	return sp.userImplementation
}

func (sp *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if sp.userRepository == nil {
		sp.userRepository = userRepository.NewRepository(sp.DBClient(ctx))
	}

	return sp.userRepository
}

func (sp *serviceProvider) UserService(ctx context.Context) service.UserService {
	if sp.userService == nil {
		sp.userService = userService.NewService(sp.UserRepository(ctx))
	}

	return sp.userService
}
