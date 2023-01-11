package main

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/wspectra/ToDoApp/internal/handler"
	"github.com/wspectra/ToDoApp/internal/repository"
	"github.com/wspectra/ToDoApp/internal/server"
	"github.com/wspectra/ToDoApp/internal/service"
	"os"
	"os/signal"
	"syscall"
)

// @title           ToDoApp
// @version         1.0
// @description     This is a simple todo list app

// @host      localhost:8080
// @BasePath  /

// @securitydefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	//structure
	if err := initConfig(); err != nil {
		log.Fatal().Msg("[CONFIG]:" + err.Error())
	}

	initLogger()
	db, err := repository.NewPostgresDB()

	if err != nil {
		log.Fatal().Msg("[DATABASE]:" + err.Error())
	}

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)
	serv := server.Server{}

	//Run
	log.Info().Msg("starting api_server on port" + viper.GetString("port"))

	go func() {
		if err := serv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			log.Fatal().Msg("[SERVER]:" + err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Info().Msg("shut  ting down api_server" + viper.GetString("port"))
	if err := serv.Shutdown(context.Background()); err != nil {
		log.Error().Msg("[SERVER]: error during shutting down" + err.Error())
	}

	if err := db.Close(); err != nil {
		log.Error().Msg("[DATABASE]: error during closing connection to database" + err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func initLogger() {
	switch viper.GetString("log_level") {
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "disabled":
		zerolog.SetGlobalLevel(zerolog.Disabled)
	default:
		log.Fatal().Msg("[CONFIG]: wrong config")
	}
}
