package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func InitLogger() {
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
