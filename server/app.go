package server

import (
	//"chat/services"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

type App struct {
	httpServer *gin.Engine
}

func NewApp() *App {
	server := gin.Default()
	//connectToDB := services.NewDbConnect()
	return &App{
		server,
	}
}

func (a *App) Run(port string) error {
	err := a.httpServer.Run(port)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
