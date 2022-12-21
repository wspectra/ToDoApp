package repository

import (
	"database/sql"
)

type Authorization interface {
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
	return &Repository{}
}
