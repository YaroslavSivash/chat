package postgres

import (
	"chat/models"
	"context"
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
	newUser := &UserCreate{UserName: user.UserName}
	_, err := r.db.Model(newUser).Insert()

	return newUser.Id, err
}

type UserCreate struct {
	tableName struct{}  `pg:"users,alias:u"`
	Id        int       `pg:",pk"`
	UserName  string    `pg:",notnull"`
	CreatedAt time.Time `pg:"default:now()"`
}
