package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/wspectra/ToDoApp/internal/structure"
)

type ItemPostgres struct {
	db *sqlx.DB
}

func NewItemPostgres(db *sqlx.DB) *ItemPostgres {
	return &ItemPostgres{db: db}
}

func (i *ItemPostgres) CreateItem(listId int, input structure.Item) error {
	tx, err := i.db.Begin()
	if err != nil {
		return err
	}

	var idItem int
	query := "INSERT INTO todo_items (title, description) VALUES ($1, $2) returning id"
	row := tx.QueryRow(query, input.Title, input.Description)
	if err := row.Scan(&idItem); err != nil {
		tx.Rollback()
		return err
	}

	query = "INSERT INTO lists_items (item_id, list_id) VALUES ($1, $2)"
	if _, err := tx.Exec(query, idItem, listId); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()

}

func (i *ItemPostgres) GetItems(listId int) ([]structure.Item, error) {
	var items []structure.Item
	query := "SELECT ti.id, ti.title, ti.description, ti.done FROM todo_items ti INNER JOIN lists_items li ON ti.id = li.item_id" +
		" WHERE li.list_id = $1"
	if err := i.db.Select(&items, query, listId); err != nil {
		return items, err
	}
	return items, nil

}

func (i *ItemPostgres) GetItemById(listId int, itemId int) (structure.Item, error) {
	var item structure.Item
	query := "SELECT ti.id, ti.title, ti.description, ti.done FROM todo_items ti INNER JOIN lists_items li ON ti.id = li.item_id" +
		" WHERE li.item_id = $1 AND li.list_id = $2"
	if err := i.db.Get(&item, query, itemId, listId); err != nil {
		return item, err
	}

	return item, nil

}

func (i *ItemPostgres) DeleteItem(itemId int) error {
	query := "DELETE FROM todo_items  WHERE id = $1"
	if _, err := i.db.Exec(query, itemId); err != nil {
		return err
	}

	return nil
}

func (i *ItemPostgres) UpdateItemTitle(itemId int, input structure.UpdateItemInput) error {
	query := "UPDATE todo_items SET title = $1 WHERE id = $2"
	if _, err := i.db.Exec(query, input.Title, itemId); err != nil {
		return err
	}

	return nil

}

func (i *ItemPostgres) UpdateItemDescription(itemId int, input structure.UpdateItemInput) error {
	query := "UPDATE todo_items SET description = $1 WHERE id = $2"
	if _, err := i.db.Exec(query, input.Description, itemId); err != nil {
		return err
	}

	return nil

}

func (i *ItemPostgres) UpdateItemDone(itemId int, input structure.UpdateItemInput) error {
	query := "UPDATE todo_items SET done = $1 WHERE id = $2"
	if _, err := i.db.Exec(query, input.Done, itemId); err != nil {
		return err
	}

	return nil

}
