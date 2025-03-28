package handlers

import (
	"context"
	"github.com/DavidMovas/SpeakUp-Server/backend/models"
	v1 "github.com/DavidMovas/SpeakUp-Server/backend/shared/grpc/v1"
	"github.com/DavidMovas/SpeakUp-Server/backend/utils/ctxc"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type ChatsHandler struct {
	ctx    context.Context
	cancel context.CancelFunc

	client v1.ChatServiceClient
	stream grpc.BidiStreamingClient[v1.ConnectRequest, v1.ConnectResponse]
}

func NewChatsHandler(client v1.ChatServiceClient) *ChatsHandler {
	var handler ChatsHandler
	handler.client = client

	ctx, cancel := context.WithCancel(context.Background())
	handler.ctx = ctx
	handler.cancel = cancel

	_ = handler.connect()

	return &handler
}

func (h *ChatsHandler) CreatePrivateChat(initiatorID, memberID string) (*models.CreateChatResponse, error) {
	response, err := h.client.CreateChat(ctxc.ShortContext(), &v1.CreateChatRequest{
		Payload: &v1.CreateChatRequest_PrivateChat_{
			PrivateChat: &v1.CreateChatRequest_PrivateChat{
				InitiatorId: initiatorID,
				MemberId:    memberID,
			},
		},
	})

	if err != nil {
		return nil, err
	}

	result := &models.CreateChatResponse{
		ChatID: response.ChatId,
	}

	return result, nil
}

func (h *ChatsHandler) CreateGroupChat(name string, initiatorID string, memberIDs []string) (*models.CreateChatResponse, error) {
	response, err := h.client.CreateChat(ctxc.ShortContext(), &v1.CreateChatRequest{
		Payload: &v1.CreateChatRequest_GroupChat_{
			GroupChat: &v1.CreateChatRequest_GroupChat{
				Name:        name,
				InitiatorId: initiatorID,
				MembersIds:  memberIDs,
			},
		},
	})

	if err != nil {
		return nil, err
	}

	result := &models.CreateChatResponse{
		ChatID: response.ChatId,
	}

	return result, nil
}

func (h *ChatsHandler) JoinChat(chatID, userID string, lastReadAt time.Time) ([]*models.Message, error) {
	if h.stream == nil {
		if err := h.connect(); err != nil {
			return nil, err
		}
	}

	err := h.stream.Send(&v1.ConnectRequest{
		Payload: &v1.ConnectRequest_JoinChat_{
			JoinChat: &v1.ConnectRequest_JoinChat{
				ChatId:     chatID,
				UserId:     userID,
				LastReadAt: timestamppb.New(lastReadAt),
			},
		},
	})

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (h *ChatsHandler) Initiate(receiveFn func() []*models.Message) (postFn func(msg *models.Message) error, err error) {
	if h.stream == nil {
		if err = h.connect(); err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func (h *ChatsHandler) connect() error {
	var err error
	h.stream, err = h.client.Connect(ctxc.ShortContext())

	if err != nil {
		return err
	}

	return nil
}

func (h *ChatsHandler) Shutdown() {
	h.cancel()
}
