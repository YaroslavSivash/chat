package server

import (
	"github.com/gin-gonic/gin"

	"net/http"
)

type App struct {
	httpServer	*http.Server

}

func NewApp() *App {
return nil
}

func (a *App) Run(port string) error {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
		)
	return nil
}