package http

import (
	"chat/chat"
	"chat/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
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
	chatJson := &CreateChatDate{}
	if err := c.BindJSON(chatJson); err != nil {
		log.Error(err)
		c.String(http.StatusBadRequest, "cannot read json")
	}

	chatID, err := h.useCase.CreateChat(c, &models.ChatEntity{Name: chatJson.Name, Users: chatJson.Users})
	if err != nil {
		log.Error(err)
		c.String(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, chatID)
}

func (h *Handler) GetAllChatUserID(c *gin.Context) {
	userJson := &GetAllChats{}
	if err := c.BindJSON(userJson); err != nil {
		log.Error(err)
		c.String(http.StatusBadRequest, "cannot read json ")
	}

	userID, err := strconv.Atoi(userJson.UserID)
	fmt.Println(userJson.UserID)
	if err != nil {
		log.Error(err)
		c.String(http.StatusBadRequest, "Not correct json")
	}
	userdata, err := h.useCase.GetAllChatUserID(c, &models.UserEntity{Id: userID})
	if err != nil {
		log.Error(err)
		c.String(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, userdata)
}

type GetAllChats struct {
	UserID string `json:"user"`
}

type CreateChatDate struct {
	Name  string   `json:"name"`
	Users []string `json:"users"`
}
