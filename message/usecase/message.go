package usecase

import (
	"chat/message"
	"chat/models"
	"context"
)

type MessageUseCase struct {
	repo message.MessageRepository
}

func NewMessageUseCase(repo message.MessageRepository) *MessageUseCase {
	return &MessageUseCase{
		repo: repo,
	}
}
func (u *MessageUseCase) SendMassage(ctx context.Context, message *models.MessageEntity) (int, error) {
	return u.repo.SendMessageDB(ctx, message)
}

func (u *MessageUseCase) GetAllMessageChatID(ctx context.Context, chat *models.ChatEntity) (*[]models.MessageEntity, error) {
	return u.repo.GetAllMessageChatIDDB(ctx, chat)
}
