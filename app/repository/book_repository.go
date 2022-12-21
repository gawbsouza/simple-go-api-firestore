package repository

import "library/entity"

type BookRepository interface {
	SelectById(id string) (*entity.Book, error)
	Insert(*entity.Book) (*entity.Book, error)
	// Update(*entity.Book) error
	// Delete(id string) error
	// SelectAll() ([]*entity.Book, error)
}
