package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/wspectra/ToDoApp/internal/structure"
)

type Authorization interface {
	AddNewUser(user structure.User) error
	AuthorizeUser(user structure.SignInUser) (int, error)
}

type List interface {
	CreateList(userId int, input structure.List) error
	GetLists(userId int) ([]structure.List, error)
	GetListById(userId int, listId int) (structure.List, error)
	DeleteList(listId int) error
	UpdateListTitle(listId int, input structure.UpdateListInput) error
	UpdateListDescription(listId int, input structure.UpdateListInput) error
	UpdateListTitleAndDescription(listId int, input structure.UpdateListInput) error
}

type Item interface {
	CreateItem(listId int, input structure.Item) error
	GetItems(listId int) ([]structure.Item, error)
	GetItemById(listId int, itemId int) (structure.Item, error)
	DeleteItem(itemId int) error
	UpdateItemTitle(itemId int, input structure.UpdateItemInput) error
	UpdateItemDescription(itemId int, input structure.UpdateItemInput) error
	UpdateItemDone(itemId int, input structure.UpdateItemInput) error
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
		Item:          NewItemPostgres(db),
	}
}
