package repository

import (
	"database/sql"
	"github.com/wspectra/api_server/internal/structure"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db}
}

func (a *AuthRepository) AddNewUser(user structure.User) error {
	//if err := a.db.QueryRow("SELECT  id FROM users where usernamne = $1", user.Username); err == nil {
	//	return errors.New("[SIGN-UP]: Username already exits")
	//}

	if _, err := a.db.Exec("INSERT INTO users (name, username, password_hash) VALUES ($1, $2, $3)",
		user.Name, user.Username, user.Password); err != nil {
		return err
	}
	return nil
}

func (a *AuthRepository) AuthorizeUser(user structure.User) error {
	var id int
	row := a.db.QueryRow("select id from users where username = $1 and password_hash = $2", user.Username, user.Password)
	err := row.Scan(&id)

	if err != nil {
		return err
	}
	return nil
}
