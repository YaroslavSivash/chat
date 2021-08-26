package user

import (
	"chat/models"
	"context"
)

type UseCaseUsers interface {
	CreateUser(ctx context.Context, user *models.UserEntity) (int, error)
}
