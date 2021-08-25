package services

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/spf13/viper"
	"os"
)

func NewDbConnect() (db *pg.DB) {
	dbconfig := Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port_db"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
	}
	pgDB, _ := dial(dbconfig)

	fmt.Println("Successful connected to DB!")
	return pgDB
}
func dial(cfg Config) (*pg.DB, error) {
	dbpg := pg.Connect(&pg.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		User:     cfg.Username,
		Password: cfg.Password,
		Database: cfg.DBName,
	})

	err := dbpg.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return dbpg, nil
}

type Config struct {
	Port     string
	Host     string
	DBName   string
	Username string
	Password string
}
