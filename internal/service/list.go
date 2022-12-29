package service

import (
	"github.com/wspectra/api_server/internal/repository"
	"github.com/wspectra/api_server/internal/structure"
)

type ListService struct {
	repo repository.List
}

func NewListService(repo repository.List) *ListService {
	return &ListService{repo: repo}
}

func (l *ListService) CreateList(userId int, input structure.List) error {
	return l.repo.CreateList(userId, input)
}

func (l *ListService) GetLists(userId int) ([]structure.List, error) {
	return l.repo.GetLists(userId)
}

func (l *ListService) GetListById(userId int, listId int) (structure.List, error) {
	return l.repo.GetListById(userId, listId)
}