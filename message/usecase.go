package message

import (
	"chat/models"
	"github.com/gin-gonic/gin"
)

type UseCaseMessage interface {
	SendMassage(c gin.Context, message *models.MessageEntity) (int, error)
	GetAllMessageChatID(c gin.Context, chat *models.ChatEntity) (*[]models.MessageEntity, error)
}
