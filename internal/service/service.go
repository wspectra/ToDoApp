package service

import (
	"github.com/wspectra/ToDoApp/internal/repository"
	"github.com/wspectra/ToDoApp/internal/structure"
)

type Authorization interface {
	AddNewUser(user structure.User) error
	GetToken(input structure.SignInUser) (string, error)
	ParseToken(accessToken string) (int, error)
}

type List interface {
	CreateList(userId int, input structure.List) error
	GetLists(userId int) ([]structure.List, error)
	GetListById(userId int, listId int) (structure.List, error)
	UpdateList(listId int, input structure.UpdateListInput) error
	DeleteList(listId int) error
}

type Item interface {
	CreateItem(userId int, listId int, input structure.Item) error
	GetItems(userId int, listId int) ([]structure.Item, error)
	GetItemById(userId int, listId int, itemId int) (structure.Item, error)
	UpdateItem(userId int, listId int, itemId int, input structure.UpdateItemInput) error
	DeleteItem(userId int, listId int, itemId int) error
}

type Service struct {
	Authorization
	List
	Item
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		List:          NewListService(repos.List),
		Item:          NewItemService(repos.Item, repos.List),
	}
}
