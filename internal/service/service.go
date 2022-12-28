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
	}
}
