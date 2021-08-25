package http

import (
	"chat/message"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase message.UseCaseMessage
}

func NewHandlerMessage(useCase message.UseCaseMessage) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) SendMessage(c *gin.Context) {
	return
}

func (h *Handler) GetAllMessageChatID(c *gin.Context) {
	return
}
