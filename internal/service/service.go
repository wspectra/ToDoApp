package service

import "github.com/wspectra/api_server/internal/repository"

type Authorization interface {
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
	return &Service{}
}
