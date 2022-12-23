package repository

import (
	"database/sql"
	"github.com/wspectra/api_server/internal/structure"
)

type Authorization interface {
	AddNewUser(user structure.User) error
	AuthorizeUser(user structure.User) error
}

type List interface {
}

type Item interface {
}

type Repository struct {
	Authorization
	List
	Item
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
	}
}
