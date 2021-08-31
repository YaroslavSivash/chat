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
	//
	tempChats := &Chats{Name: chat.Name}
	//err := r.db.Model(tempChats).Where("name= ?", chat.Name).Select()
	//if err != nil {
	//	return 0, errors.New("this name is exist")
	//}

	tx, _ := r.db.Begin()

	_, err := tx.Model(tempChats).Returning("*").Insert()
	if err != nil {
		return 0, errors.New("cannot add chat")
	}

	// Make sure to close transaction if something goes wrong.

	if err := addUsersChats(tx, &intSlice, tempChats.Id); err != nil {
		// Rollback on error.
		_ = tx.Rollback()
		return 0, errors.New("error")
	}

	// Commit on success.
	if err := tx.Commit(); err != nil {
		log.Error(err)
	}
	return tempChats.Id, nil
}
func addUsersChats(tx *pg.Tx, users *[]int, ChatID int) error {
	for _, user := range *users {
		if err := tx.Model(UsersChats{UserID: user, ChatID: ChatID}).Select(); err != nil {
			log.Error(err)
			return nil
		}
	}
	return nil
}
func (r *ChatRepository) GetAllChatUserIDDB(ctx context.Context, user *models.UserEntity) (*[]models.ChatEntity, error) {
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
	Users     []string  `pg:"users,notnull"`
	CreatedAt time.Time `pg:"default:now()"`
}
