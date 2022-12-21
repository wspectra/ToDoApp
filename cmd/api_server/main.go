package main

import (
	"github.com/spf13/viper"
	"github.com/wspectra/api_server/internal/handler"
	"github.com/wspectra/api_server/internal/repository"
	"github.com/wspectra/api_server/internal/server"
	"github.com/wspectra/api_server/internal/service"
	"log"
)

func main() {
	//init
	if err := initConfig(); err != nil {
		log.Fatalf("[CONFIG]: %s", err.Error())
	}

	db, err := repository.NewPostgresDB()

	if err != nil {
		log.Fatalf("[DATABASE]: %s", err.Error())
	}

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)
	serv := server.Server{}

	//Run
	if err := serv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("[SERVER]: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
