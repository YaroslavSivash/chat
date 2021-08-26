package chat

import (
	"chat/models"
	"context"
)

type ChatRepository interface {
	CreateChatDB(ctx context.Context, chat *models.ChatEntity) (int, error)
	GetAllChatUserIDDB(ctx context.Context, user *models.UserEntity) (*[]models.ChatEntity, error)
}
