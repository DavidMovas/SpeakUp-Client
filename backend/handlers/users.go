package handlers

import (
	"context"
	"github.com/DavidMovas/SpeakUp-Server/backend/models"
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

func (h *UsersHandler) RegisterUser(fullName, email, password string) (*models.RegisterUserResponse, error) {
	response, err := h.client.Register(context.Background(), &v1.RegisterRequest{
		FullName: fullName,
		Email:    email,
		Password: password,
	})

	if err != nil {
		return nil, err
	}

	result := &models.RegisterUserResponse{
		AccessToken: response.AccessToken,
		User: &models.User{
			ID:        response.User.Id,
			Email:     response.User.Email,
			Username:  response.User.Username,
			FullName:  response.User.FullName,
			CreatedAt: response.User.CreatedAt.String(),
		},
	}

	return result, nil
}

func (h *UsersHandler) LoginUser(email, password string) (*models.LoginUserResponse, error) {
	response, err := h.client.Login(context.Background(), &v1.LoginRequest{
		Email:    email,
		Password: password,
	})

	if err != nil {
		return nil, err
	}

	result := &models.LoginUserResponse{
		AccessToken: response.AccessToken,
		User: &models.User{
			ID:          response.User.Id,
			Email:       response.User.Email,
			Username:    response.User.Username,
			FullName:    response.User.FullName,
			AvatarURL:   response.User.AvatarUrl,
			Bio:         response.User.Bio,
			LastLoginAt: response.User.LastLoginAt.String(),
			CreatedAt:   response.User.CreatedAt.String(),
			UpdatedAt:   response.User.UpdatedAt.String(),
		},
	}

	return result, nil
}
