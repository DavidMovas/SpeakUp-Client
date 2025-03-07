package internal

import (
	"context"

	"github.com/DavidMovas/SpeakUp-Client/internal/config"
	"google.golang.org/grpc"
)

type Server struct {
	grpcServer *grpc.Server
}

func NewServer(ctx context.Context, cfg config.Config) (*Server, error) {

}

func (s *Server) Start() error {

	return nil
}

func (s *Server) Stop(ctx context.Context) error {

	return nil
}

func (s *Server) close() error {

	return nil
}
