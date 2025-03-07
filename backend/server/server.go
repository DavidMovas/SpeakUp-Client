package server

import (
	"context"
	"fmt"
	"github.com/DavidMovas/SpeakUp-Server/backend/config"
	"github.com/DavidMovas/SpeakUp-Server/backend/handlers"
	v1 "github.com/DavidMovas/SpeakUp-Server/backend/shared/grpc/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	UsersHandler *handlers.UsersHandler
	cfg          *config.Config
	closers      []func() error
}

func NewServer(ctx context.Context, cfg *config.Config) (*Server, error) {
	var closers []func() error

	conn, err := grpc.NewClient(cfg.APIServerTCPURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("could not create grpc connection: %w", err)
	}

	closers = append(closers, conn.Close)

	v1.NewUsersServiceClient(conn)

	return &Server{
		cfg:     cfg,
		closers: closers,
	}, nil
}

func (s *Server) Close() error {
	return
}
