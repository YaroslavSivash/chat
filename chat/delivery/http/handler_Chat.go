package http

import (
	"chat/chat"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase chat.UseCaseChat
}

func NewHandlerChat(useCase chat.UseCaseChat) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) CreateChat(c *gin.Context) {
	return
}
func (h *Handler) GetAllChatUserID(c *gin.Context) {
	return
}
