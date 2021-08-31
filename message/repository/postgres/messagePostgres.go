package postgres

import (
	"chat/models"
	"context"
	"errors"
	"github.com/go-pg/pg/v10"
	"github.com/labstack/gommon/log"
	"strconv"
	"time"
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
	intChat, err := strconv.Atoi(message.Chat)
	intUser, err := strconv.Atoi(message.Author)
	mes := &Message{Chat: intChat, Author: intUser, Text: message.Text}
	_, err = r.db.Model(mes).Insert()
	if err != nil {
		log.Error(err)
	}
	return mes.Id, err
}

func (r *MessageRepository) GetAllMessageChatIDDB(ctx context.Context, chat *models.ChatEntity) (*[]models.MessageEntity, error) {

	mes := &[]Message{}

	err := r.db.Model(mes).OrderExpr("created_at DESC").Where("chat_id = ?", chat.Id).Select()
	if err != nil {
		log.Error(err)
		return nil, errors.New("something wrong")
	}

	allMes := []models.MessageEntity{}
	for _, message := range *mes {
		chatString := strconv.Itoa(message.Chat)
		authorString := strconv.Itoa(message.Author)
		allMes = append(allMes, models.MessageEntity{
			Id:        message.Id,
			Chat:      chatString,
			Author:    authorString,
			Text:      message.Text,
			CreatedAt: message.CreatedAt,
		})
	}

	return &allMes, nil
}

type Message struct {
	tableName struct{}  `pg:"message,alias:message"`
	Id        int       `pg:"id,pk"`
	Chat      int       `pg:"chat_id,notnull"`
	Author    int       `pg:"author_id,notnull"`
	Text      string    `pg:"text,notnull"`
	CreatedAt time.Time `pg:"default:now()"`
}
