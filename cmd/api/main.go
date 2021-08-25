package main

import (
	"chat/config"
	"chat/server"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	app := server.NewApp()

	if err := app.Run(":" + viper.GetString("port")); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
