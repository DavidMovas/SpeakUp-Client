package server

import (
	"context"
	"fmt"
	"github.com/DavidMovas/SpeakUp-Server/backend/config"
	"github.com/DavidMovas/SpeakUp-Server/backend/handlers"
	v1 "github.com/DavidMovas/SpeakUp-Server/backend/shared/grpc/v1"
	"github.com/DavidMovas/SpeakUp-Server/backend/utils/helpers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	UsersHandler *handlers.UsersHandler
	ChatsHandler *handlers.ChatsHandler
	cfg          *config.Config
	closers      []func() error
}

func NewServer(_ context.Context, cfg *config.Config) (*Server, error) {
	var closers []func() error

	conn, err := grpc.NewClient(cfg.APIServerTCPURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("could not create grpc connection: %w", err)
	}

	closers = append(closers, conn.Close)

	usersClient := v1.NewUsersServiceClient(conn)
	chatsClient := v1.NewChatServiceClient(conn)

	usersHandler := handlers.NewUsersHandler(usersClient)
	chatsHandler := handlers.NewChatsHandler(chatsClient)

	return &Server{
		UsersHandler: usersHandler,
		ChatsHandler: chatsHandler,
		cfg:          cfg,
		closers:      closers,
	}, nil
}

func (s *Server) Stop() error {
	return helpers.WithClosers(s.closers, nil)
}
