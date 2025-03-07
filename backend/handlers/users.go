package handlers

import (
	"context"
	v1 "github.com/DavidMovas/SpeakUp-Server/backend/shared/grpc/v1"
)

type UsersHandler struct {
	client v1.UsersServiceClient
}

func NewUsersHandler(client v1.UsersServiceClient) *UsersHandler {
	return &UsersHandler{
		client: client,
	}
}

func (h *UsersHandler) RegisterUser(fullName, email, password string) (*v1.RegisterResponse, error) {
	response, err := h.client.Register(context.Background(), &v1.RegisterRequest{
		FullName: fullName,
		Email:    email,
		Password: password,
	})

	return response, err
}
