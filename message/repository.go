package message

import (
	"chat/models"
	"context"
)

type MessageRepository interface {
	SendMessageDB(ctx context.Context, message *models.MessageEntity) (int, error)
	GetAllMessageChatIDDB(ctx context.Context, chat *models.ChatEntity) (*[]models.MessageEntity, error)
}
