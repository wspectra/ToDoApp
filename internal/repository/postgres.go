package repository

import (
	"database/sql"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

const (
	usersTable      = "users"
	ListsTable      = "todo_lists"
	usersListsTable = "users_lists"
	ItemsTable      = "todo_items"
	listsItemsTable = "lists_items"
)

func NewPostgresDB() (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s dbname=%s sslmode=%s password=%s user=%s port=%s",
		viper.GetString("database.host"),
		viper.GetString("database.db_name"),
		viper.GetString("database.ssl_mode"),
		viper.GetString("database.password"),
		viper.GetString("database.user"),
		viper.GetString("database.port"))

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Info().Msg("[DATABASE]: SUCCESSFULLY CONNECTED TO DATABASE")
	return db, nil
}
