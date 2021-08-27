package postgres

import (
	"chat/models"
	"context"
	"github.com/go-pg/pg/v10"
)

type MessageRepository struct {
	db *pg.DB
}

func NewMessageRepository(db *pg.DB) *MessageRepository {
	return &MessageRepository{
		db: db,
	}
}

func (r *MessageRepository) SendMessageDB(ctx context.Context, message *models.MessageEntity) (int, error) {
	return 0, nil
}

func (r *MessageRepository) GetAllMessageChatIDDB(ctx context.Context, chat *models.ChatEntity) (*[]models.MessageEntity, error) {
	return nil, nil
}
