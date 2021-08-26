package http

import (
	"chat/models"
	"chat/user"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"net/http"
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
	userBody := &CreateUserDate{}
	if err := c.BindJSON(userBody); err != nil {
		log.Error(err)
		c.String(http.StatusBadRequest, "Cannot read json")
	}
	userID, err := h.useCase.CreateUser(c, &models.UserEntity{UserName: userBody.UserName})
	if err != nil {
		log.Error(err)
		c.String(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, userID)
}

type CreateUserDate struct {
	UserName string `json:"username"`
}
