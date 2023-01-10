package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/wspectra/api_server/internal/structure"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db}
}

func (a *AuthRepository) AddNewUser(user structure.User) error {

	if _, err := a.db.Exec("INSERT INTO users (name, username, password_hash) VALUES ($1, $2, $3)",
		user.Name, user.Username, user.Password); err != nil {
		return err
	}
	return nil
}

func (a *AuthRepository) AuthorizeUser(user structure.SignInUser) (int, error) {
	var id int
	row := a.db.QueryRow("select id from users where username = $1 and password_hash = $2", user.Username, user.Password)
	err := row.Scan(&id)

	if err != nil {
		return -1, err
	}
	return id, nil
}
