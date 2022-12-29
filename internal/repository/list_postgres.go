package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/wspectra/api_server/internal/structure"
)

type ListPostgres struct {
	db *sqlx.DB
}

func NewListPostgres(db *sqlx.DB) *ListPostgres {
	return &ListPostgres{db: db}
}

func (l *ListPostgres) CreateList(userId int, input structure.List) error {
	tx, err := l.db.Begin()
	if err != nil {
		return err
	}

	var idList int
	query := "INSERT INTO todo_lists (title, description) VALUES ($1, $2) returning id"
	row := tx.QueryRow(query, input.Title, input.Description)
	if err := row.Scan(&idList); err != nil {
		tx.Rollback()
		return err
	}

	query = "INSERT INTO users_lists (user_id, list_id) VALUES ($1, $2)"
	if _, err := tx.Exec(query, userId, idList); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (l *ListPostgres) GetLists(userId int) ([]structure.List, error) {
	var lists []structure.List
	query := "SELECT tl.id, tl.title, tl.description FROM todo_lists tl INNER JOIN users_lists ul ON tl.id = ul.list_id" +
		" WHERE ul.user_id = $1"
	if err := l.db.Select(&lists, query, userId); err != nil {
		return lists, err
	}

	return lists, nil
}

func (l *ListPostgres) GetListById(userId int, listId int) (structure.List, error) {
	var lists structure.List
	query := "SELECT tl.id, tl.title, tl.description FROM todo_lists tl INNER JOIN users_lists ul ON tl.id = ul.list_id" +
		" WHERE ul.user_id = $1 AND ul.list_id = $2"
	if err := l.db.Get(&lists, query, userId, listId); err != nil {
		return lists, err
	}

	return lists, nil

}