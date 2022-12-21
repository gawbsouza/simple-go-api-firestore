package inmemory

import (
	"errors"
	"library/entity"
	"library/repository"

	"github.com/google/uuid"
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

func (r *rep) Insert(b *entity.Book) (*entity.Book, error) {

	newID := uuid.NewString()

	if b.ID != "" {
		newID = b.ID
	}

	r.memory[newID] = b
	b.ID = newID

	return b, nil
}
