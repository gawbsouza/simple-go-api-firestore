package usecases

import (
	"errors"
	"library/entity"

	"github.com/go-playground/validator/v10"
)

type CreateBookUseCase interface {
	Execute(*CreateBookUseCaseInputDTO) (*CreateBookUseCaseOutputDTO, error)
}

type CreateBookUseCaseInputDTO struct {
	Title   string   `validate:"required"`
	Authors []string `validate:"min=1"`
	Year    int      `validate:"gt=0"`
	Edition int      `validate:"gt=0"`
	Pages   int      `validate:"gt=0"`
}

type CreateBookUseCaseOutputDTO struct {
	ID      string
	Title   string
	Authors []string
	Year    int
	Edition int
	Pages   int
}

type CreateBookUseCaseRepository interface {
	Insert(*entity.Book) (*entity.Book, error)
}

type createNewBook struct {
	repository CreateBookUseCaseRepository
}

func NewCreateBookUseCase(rep CreateBookUseCaseRepository) CreateBookUseCase {
	return &createNewBook{rep}
}

func (u *createNewBook) Execute(input *CreateBookUseCaseInputDTO) (*CreateBookUseCaseOutputDTO, error) {

	validate := validator.New()

	err := validate.Struct(input)
	if err != nil {
		return nil, err
	}

	toCreate := &entity.Book{
		ID:      "",
		Title:   input.Title,
		Authors: input.Authors,
		Year:    input.Year,
		Edition: input.Edition,
		Pages:   input.Pages,
	}

	book, err := u.repository.Insert(toCreate)
	if err != nil {
		return nil, errors.New("Error when inserting book into repository")
	}

	output := &CreateBookUseCaseOutputDTO{
		ID:      book.ID,
		Title:   book.Title,
		Authors: book.Authors,
		Year:    book.Year,
		Edition: book.Edition,
		Pages:   book.Pages,
	}

	return output, nil
}
