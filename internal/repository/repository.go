package repository

type Authorization interface {
}

type List interface {
}

type Item interface {
}

type Repository struct {
	Authorization
	List
	Item
}

func NewRepository() *Repository {
	return &Repository{}
}
