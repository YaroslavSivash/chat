package http

import (
	"chat/message"
	"github.com/gin-gonic/gin"
)

func RegisterHttpEndPointsMessage(c *gin.Engine, ms message.UseCaseMessage) {
	h := NewHandlerMessage(ms)

	message := c.Group("/message")
	{
		message.POST("/add", h.SendMessage)
		message.POST("/get", h.GetAllMessageChatID)
	}
}
