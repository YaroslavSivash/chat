package postgres

import (
	"chat/models"
	"context"
	"github.com/go-pg/pg/v10"
)

type ChatRepository struct {
	db *pg.DB
}

func NewChatRepository(db *pg.DB) *ChatRepository {
	return &ChatRepository{
		db: db,
	}
}
func (r *ChatRepository) CreateChatDB(ctx context.Context, chat *models.ChatEntity) (int, error) {
	return 0, nil
}
func (r *ChatRepository) GetAllChatUserIDDB(ctx context.Context, user *models.UserEntity) (*[]models.ChatEntity, error) {
	return nil, nil
}
