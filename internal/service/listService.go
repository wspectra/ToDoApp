package service

import (
	"github.com/wspectra/ToDoApp/internal/repository"
	"github.com/wspectra/ToDoApp/internal/structure"
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

func (l *ListService) UpdateList(listId int, input structure.UpdateListInput) error {
	if input.Title != nil && input.Description == nil {
		return l.repo.UpdateListTitle(listId, input)
	}
	if input.Title == nil && input.Description != nil {
		return l.repo.UpdateListDescription(listId, input)
	}
	if input.Title != nil && input.Description != nil {
		return l.repo.UpdateListTitleAndDescription(listId, input)
	}
	return nil
}

func (l *ListService) DeleteList(listId int) error {
	return l.repo.DeleteList(listId)
}
