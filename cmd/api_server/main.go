package main

import (
	"github.com/wspectra/api_server/internal/handler"
	"github.com/wspectra/api_server/internal/repository"
	"github.com/wspectra/api_server/internal/server"
	"github.com/wspectra/api_server/internal/service"
	"log"
)

func main() {
	//init
	repositories := repository.NewRepository()
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)
	server := server.Server{}

	//Run
	if err := server.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("[SERVER]: %s", err.Error())
	}
}
