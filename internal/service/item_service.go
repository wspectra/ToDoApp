package service

import (
	"errors"
	"github.com/wspectra/api_server/internal/repository"
	"github.com/wspectra/api_server/internal/structure"
)

type ItemService struct {
	repo     repository.Item
	listRepo repository.List
}

func NewItemService(repo repository.Item, repoList repository.List) *ItemService {
	return &ItemService{repo: repo, listRepo: repoList}
}

func (i *ItemService) CreateItem(userId int, listId int, input structure.Item) error {
	if _, err := i.listRepo.GetListById(userId, listId); err != nil {
		return errors.New("list does not exist or does not belong to user")
	}
	return i.repo.CreateItem(listId, input)
}

func (i *ItemService) GetItems(userId int, listId int) ([]structure.Item, error) {
	if _, err := i.listRepo.GetListById(userId, listId); err != nil {
		return nil, errors.New("list does not exist or does not belong to user")
	}
	return i.repo.GetItems(listId)
}

func (i *ItemService) GetItemById(userId int, listId int, itemId int) (structure.Item, error) {
	if _, err := i.listRepo.GetListById(userId, listId); err != nil {
		return structure.Item{}, errors.New("list does not exist or does not belong to user")
	}
	return i.repo.GetItemById(listId, itemId)
}

func (i *ItemService) UpdateItem(userId int, listId int, itemId int, input structure.UpdateItemInput) error {
	if _, err := i.listRepo.GetListById(userId, listId); err != nil {
		return errors.New("list does not exist or does not belong to user")
	}

	if input.Title != nil {
		if err := i.repo.UpdateItemTitle(itemId, input); err != nil {
			return err
		}
	}
	if input.Description != nil {
		if err := i.repo.UpdateItemDescription(itemId, input); err != nil {
			return err
		}
	}
	if input.Done != nil {
		if err := i.repo.UpdateItemDone(itemId, input); err != nil {
			return err
		}
	}
	return nil
}

func (i *ItemService) DeleteItem(userId int, listId int, itemId int) error {
	if _, err := i.listRepo.GetListById(userId, listId); err != nil {
		return err
	}
	return i.repo.DeleteItem(itemId)
}
