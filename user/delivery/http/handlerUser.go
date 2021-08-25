package http

import (
	"chat/user"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase user.UseCaseUsers
}

func NewHandlerUser(useCase user.UseCaseUsers) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) CreateUser(c *gin.Context) {
	//userBody := &models.UserEntity{}
	return
}
