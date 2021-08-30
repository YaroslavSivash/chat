package postgres

import (
	"chat/models"
	"context"
	"github.com/labstack/gommon/log"

	"github.com/go-pg/pg/v10"
	"time"
)

type UserRepository struct {
	db *pg.DB
}

func NewUserRepository(db *pg.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUserDB(ctx context.Context, user *models.UserEntity) (int, error) {
	newUser := &Users{UserName: user.UserName}
	_, err := r.db.Model(newUser).Insert()
	if err != nil {
		log.Error(err)
	}

	return newUser.Id, err
}

type Users struct {
	tableName struct{}  `pg:"users,alias:users"`
	Id        int       `pg:"id,pk"`
	UserName  string    `pg:"username,notnull"`
	CreatedAt time.Time `pg:"default:now()"`
}
