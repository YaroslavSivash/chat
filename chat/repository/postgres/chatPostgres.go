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
	intSlice := make([]int, 0, len(chat.Users))

	for i, user := range chat.Users {
		temp, err := strconv.Atoi(user)
		if err != nil {
			log.Error(err)
			return 0, errors.New("cannot convert string slice :(")
		}
		intSlice[i] = temp
	}
	tempChats := &Chats{}
	err := r.db.Model(tempChats).Where("name= ?", tempChats.Name)
	if err != nil {
		return 0, errors.New("this name is exist")
	}

	tx, err := r.db.Begin()

	// Make sure to close transaction if something goes wrong.
	defer tx.Close()

	if err := doSomething(ctx, tx); err != nil {
		// Rollback on error.
		_ = tx.Rollback()
		return nil, nil
	}

	// Commit on success.
	if err := tx.Commit(); err != nil {
		panic(err)
	}
	return 0, nil
}
func (r *ChatRepository) GetAllChatUserIDDB(ctx context.Context, user *models.UserEntity) (*[]models.ChatEntity, error) {
	return nil, nil
}

type UsersChats struct {
	tableName struct{} `pg:"userschats,alias:userschats"`
	Id        int      `pg:"id,pk"`
	UserID    int      `pg:"user_id,notnull"`
	ChatID    int      `pg:"chat_id,notnull"`
}

type Chats struct {
	tableName struct{}  `pg:"chats,alias:chats"`
	Name      string    `pg:"name,notnull"`
	CreatedAt time.Time `pg:"default:now()"`
}
