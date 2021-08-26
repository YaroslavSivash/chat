package user

import (
	"chat/models"
	"context"
)

type UserReository interface {
	CreateUserDB(ctx context.Context, user *models.UserEntity) (int, error)
}
