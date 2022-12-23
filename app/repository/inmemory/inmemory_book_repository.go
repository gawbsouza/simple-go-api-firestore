package repository

import (
	"errors"
	"library/entity"
	"library/repository"

	"github.com/google/uuid"
)

type rep struct {
	memory map[string]entity.Book
}

func NewInMemoryBookRepository() repository.BookRepository {
	return &rep{memory: make(map[string]entity.Book)}
}

func (r *rep) SelectById(id string) (*entity.Book, error) {

	book, found := r.memory[id]

	if found {
		return &book, nil
	}

	return nil, errors.New("Book not found")
}

func (r *rep) Insert(b *entity.Book) (*entity.Book, error) {

	if b.ID == "" {
		b.ID = uuid.NewString()
	}

	r.memory[b.ID] = *b

	return b, nil
}

func (r *rep) SelectAll() ([]entity.Book, error) {

	var output []entity.Book

	for _, book := range r.memory {
		output = append(output, book)
	}

	return output, nil
}

func (r *rep) Delete(id string) error {

	if _, ok := r.memory[id]; !ok {
		return errors.New("Book not found")
	}

	delete(r.memory, id)
	return nil

}

func (r *rep) Update(b *entity.Book) error {

	if _, ok := r.memory[b.ID]; ok {
		return errors.New("Book not found")
	}

	r.memory[b.ID] = *b
	return nil
}
