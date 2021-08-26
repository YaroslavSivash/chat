package message

import (
	"chat/models"
	"context"
)

type UseCaseMessage interface {
	SendMassage(ctx context.Context, message *models.MessageEntity) (int, error)
	GetAllMessageChatID(ctx context.Context, chat *models.ChatEntity) (*[]models.MessageEntity, error)
}
