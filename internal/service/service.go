package service

import (
	"github.com/wspectra/api_server/internal/repository"
	"github.com/wspectra/api_server/internal/structure"
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
}

type Item interface {
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
	}
}
