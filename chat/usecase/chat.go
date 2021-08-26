package usecase

import (
	"chat/chat"
	"chat/models"
	"context"
)

type ChatUseCase struct {
	repo chat.ChatRepository
}

func NewChatUseCase(repo chat.ChatRepository) *ChatUseCase {
	return &ChatUseCase{
		repo: repo,
	}
}

func (u *ChatUseCase) CreateChat(ctx context.Context, chat *models.ChatEntity) (int, error) {
	return u.repo.CreateChatDB(ctx, chat)
}
func (u *ChatUseCase) GetAllChatUserIDDB(ctx context.Context, user *models.UserEntity) (*[]models.ChatEntity, error) {
	return u.repo.GetAllChatUserIDDB(ctx, user)
}
