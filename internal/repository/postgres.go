package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/wspectra/ToDoApp/internal/pkg/utils"
	"time"

	_ "github.com/lib/pq"
)

const (
	usersTable      = "users"
	ListsTable      = "todo_lists"
	usersListsTable = "users_lists"
	ItemsTable      = "todo_items"
	listsItemsTable = "lists_items"
)

func NewPostgresDB() (*sqlx.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s dbname=%s sslmode=%s password=%s user=%s port=%s",
		viper.GetString("database.host"),
		viper.GetString("database.db_name"),
		viper.GetString("database.ssl_mode"),
		viper.GetString("database.password"),
		viper.GetString("database.user"),
		viper.GetString("database.port"))

	log.Log().Msg(dataSourceName)

	var db *sqlx.DB

	err := utils.DoWithTries(func() error {
		var err error
		db, err = sqlx.Open("postgres", dataSourceName)
		if err != nil {
			return err
		}
		err = db.Ping()
		if err != nil {
			return err
		}
		return nil
	}, 10, 5*time.Second)

	if err != nil {
		return nil, err
	}

	log.Info().Msg("[DATABASE]: SUCCESSFULLY CONNECTED TO DATABASE")
	return db, nil
}
