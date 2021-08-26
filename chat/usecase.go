package chat

import (
	"chat/models"
	"context"
)

type UseCaseChat interface {
	CreateChat(ctx context.Context, chat *models.ChatEntity) (int, error)
	GetAllChatUserID(ctx context.Context, user *models.UserEntity) (*[]models.ChatEntity, error)
}
