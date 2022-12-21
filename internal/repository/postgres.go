package repository

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

func NewPostgresDB() (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s dbname=%s sslmode=%s password=%s user=%s",
		viper.GetString("database.host"),
		viper.GetString("database.db_name"),
		viper.GetString("database.ssl_mode"),
		viper.GetString("database.password"),
		viper.GetString("database.username"))

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
