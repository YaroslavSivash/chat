package usecase

import (
	"chat/models"
	"chat/user"
	"context"
)

type UserUseCase struct {
	repo user.UserReository
}

func NewUserUseCase(repo user.UserReository) *UserUseCase {
	return &UserUseCase{
		repo: repo,
	}
}

func (u *UserUseCase) CreateUser(ctx context.Context, user *models.UserEntity) (int, error) {
	return u.repo.CreateUserDB(ctx, user)
}
