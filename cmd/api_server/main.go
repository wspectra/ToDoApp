package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/wspectra/api_server/internal/handler"
	"github.com/wspectra/api_server/internal/repository"
	"github.com/wspectra/api_server/internal/server"
	"github.com/wspectra/api_server/internal/service"
)

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
	if err := serv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatal().Msg("[SERVER]:" + err.Error())
	}
	log.Info().Msg("starting api_server on port" + viper.GetString("port"))
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
