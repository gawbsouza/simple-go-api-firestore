package inmemory

import (
	"errors"
	"library/entity"
	"library/repository"
)

type rep struct {
	memory map[string]*entity.Book
}

func NewInMemoryBookRepository() repository.BookRepository {
	return &rep{memory: make(map[string]*entity.Book)}
}

func (r *rep) SelectById(id string) (*entity.Book, error) {

	book := r.memory[id]

	if book != nil {
		return book, nil
	}

	return nil, errors.New("Book not found")
}
