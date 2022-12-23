package service

import (
	"github.com/wspectra/api_server/internal/pkg/utils"
	"github.com/wspectra/api_server/internal/repository"
	"github.com/wspectra/api_server/internal/structure"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) AddNewUser(user structure.User) error {
	user.Password = utils.GeneratePasswordHash(user.Password)
	return a.repo.AddNewUser(user)
}

func (a *AuthService) AuthorizeUser(user structure.User) error {
	user.Password = utils.GeneratePasswordHash(user.Password)
	return a.repo.AuthorizeUser(user)
}
