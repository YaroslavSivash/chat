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

type ChatRepository struct {
	db *pg.DB
}

func NewChatRepository(db *pg.DB) *ChatRepository {
	return &ChatRepository{
		db: db,
	}
}
func (r *ChatRepository) CreateChatDB(ctx context.Context, chat *models.ChatEntity) (int, error) {

	intSlice := make([]int, len(chat.Users), len(chat.Users))

	for i, user := range chat.Users {
		temp, err := strconv.Atoi(user)
		if err != nil {
			log.Error(err)
			return 0, errors.New("cannot convert string slice :(")
		}
		intSlice[i] = temp
	}

	tx, err := r.db.Begin()
	if err != nil {
		log.Error(err)
		return 0, err
	}

	defer tx.Close()

	tempChats := &Chats{Name: chat.Name}
	log.Info(chat.Name)
	_, err = r.db.Model(tempChats).Insert()
	if err != nil {
		log.Error(err)
		return 0, err
	}

	if err := AddUserChats(tx, intSlice, tempChats.Id); err != nil {
		// Rollback on error.
		_ = tx.Rollback()
		log.Error(err)
		return 0, err
	}

	// Commit on success.
	if err := tx.Commit(); err != nil {
		log.Error(err)
		return 0, err
	}

	return tempChats.Id, err
}

func AddUserChats(tx *pg.Tx, userID []int, chatID int) error {
	for _, user := range userID {
		log.Info(chatID)
		if _, err := tx.Model(&UsersChats{UserID: user, ChatID: chatID}).Insert(); err != nil {
			log.Error(err)
			return nil
		}
	}
	return nil
}
func (r *ChatRepository) GetAllChatUserIDDB(ctx context.Context, user *models.UserEntity) (*[]models.ChatEntity, error) {

	usrChats := &UsersChats{}

	log.Info(user.Id)

	err := r.db.Model(usrChats).Where("user_id = ?", user.Id).Select("chat_id")
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return nil, nil
}

type UsersChats struct {
	tableName struct{} `pg:"users_chats,alias:users_chats"`
	Id        int      `pg:"id,pk"`
	UserID    int      `pg:"user_id,notnull"`
	ChatID    int      `pg:"chat_id,notnull"`
}

type Chats struct {
	tableName struct{}  `pg:"chats,alias:chats"`
	Id        int       `pg:"id,pk"`
	Name      string    `pg:"name,notnull"`
	CreatedAt time.Time `pg:"default:now()"`
}
