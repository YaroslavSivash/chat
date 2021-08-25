package user

import (
	"chat/models"
	"github.com/gin-gonic/gin"
)

type UseCaseUsers interface {
	CreateUser(c gin.Context, user *models.UserEntity) (int, error)
}
