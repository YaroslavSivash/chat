package chat

import (
	"chat/models"
	"github.com/gin-gonic/gin"
)

type UseCaseChat interface {
	CreateChat(c gin.Context, chat *models.ChatEntity) (int, error)
	GetAllChatUserID(c gin.Context, user *models.UserEntity) (*[]models.ChatEntity, error)
}
