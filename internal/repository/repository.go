package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/wspectra/api_server/internal/structure"
)

type Authorization interface {
	AddNewUser(user structure.User) error
	AuthorizeUser(user structure.SignInUser) (int, error)
}

type List interface {
	CreateList(userId int, input structure.List) error
	GetLists(userId int) ([]structure.List, error)
	GetListById(userId int, listId int) (structure.List, error)
}

type Item interface {
}

type Repository struct {
	Authorization
	List
	Item
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		List:          NewListPostgres(db),
	}
}
