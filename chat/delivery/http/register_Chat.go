package http

import (
	"chat/chat"
	"github.com/gin-gonic/gin"
)

func RegisterHttpEndPointsChat(c *gin.Engine, ch chat.UseCaseChat) {
	h := NewHandlerChat(ch)

	chats := c.Group("/chats")
	{
		chats.POST("/add", h.CreateChat)
		chats.POST("/get", h.GetAllChatUserID)
	}
}
