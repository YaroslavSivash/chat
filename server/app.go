package server

import (
	"chat/chat"
	"chat/message"

	//"chat/chat"
	chatHttp "chat/chat/delivery/http"
	chatRepo "chat/chat/repository/postgres"
	chatUseCase "chat/chat/usecase"
	//"chat/message"
	messageHttp "chat/message/delivery/http"
	messageRepo "chat/message/repository/postgres"
	messageUseCase "chat/message/usecase"
	"chat/services"
	"chat/user"
	userHttp "chat/user/delivery/http"
	userRepo "chat/user/repository/postgres"
	userUseCase "chat/user/usecase"

	//"chat/services"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

type App struct {
	httpServer *gin.Engine
	UserUC     user.UseCaseUsers
	MessageUC  message.UseCaseMessage
	ChatUC     chat.UseCaseChat
}

func NewApp() *App {
	server := gin.Default()
	connectToPostgres := services.NewDbConnect()
	repoUser := userRepo.NewUserRepository(connectToPostgres)
	repoMessage := messageRepo.NewMessageRepository(connectToPostgres)
	repoChat := chatRepo.NewChatRepository(connectToPostgres)

	return &App{
		server,
		userUseCase.NewUserUseCase(repoUser),
		messageUseCase.NewMessageUseCase(repoMessage),
		chatUseCase.NewChatUseCase(repoChat),
	}
}

func (a *App) Run(port string) error {

	userHttp.RegisterHttpEndPointsUser(a.httpServer, a.UserUC)
	messageHttp.RegisterHttpEndPointsMessage(a.httpServer, a.MessageUC)
	chatHttp.RegisterHttpEndPointsChat(a.httpServer, a.ChatUC)
	err := a.httpServer.Run(port)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
