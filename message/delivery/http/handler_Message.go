package http

import (
	"chat/message"
	"chat/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
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
	messageBody := &CreateMessage{}
	if err := c.BindJSON(messageBody); err != nil {
		log.Error(err)
		c.String(http.StatusBadRequest, "cannot read json")
	}

	_, err := strconv.Atoi(messageBody.UserID)
	fmt.Println(messageBody.UserID)
	if err != nil {
		log.Error(err)
		c.String(http.StatusBadRequest, "Not correct json")
	}
	_, err = strconv.Atoi(messageBody.ChatID)
	fmt.Println(messageBody.ChatID)
	if err != nil {
		log.Error(err)
		c.String(http.StatusBadRequest, "Not correct json")
	}

	messageID, err := h.useCase.SendMassage(c, &models.MessageEntity{Chat: messageBody.ChatID, Author: messageBody.UserID, Text: messageBody.Text})
	if err != nil {
		log.Error(err)
		c.String(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, messageID)
}

func (h *Handler) GetAllMessageChatID(c *gin.Context) {
	messageJson := &CreateMessage{}
	if err := c.BindJSON(messageJson); err != nil {
		log.Error(err)
		c.String(http.StatusBadRequest, "cannot read json ")
	}

	chatID, err := strconv.Atoi(messageJson.ChatID)
	fmt.Println(messageJson.ChatID)
	if err != nil {
		log.Error(err)
		c.String(http.StatusBadRequest, "Not correct json")
	}
	messageData, err := h.useCase.GetAllMessageChatID(c, &models.ChatEntity{Id: chatID})
	if err != nil {
		log.Error(err)
		c.String(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, messageData)
}

type CreateMessage struct {
	ChatID string `json:"chat"`
	UserID string `json:"author"`
	Text   string `json:"text"`
}
