package usecases

import (
	"errors"
	"library/repository"

	"github.com/go-playground/validator/v10"
)

type GetBookByID interface {
	Execute(*GetBookByIDInputDTO) (*GetBookByIDOutputDTO, error)
}

type GetBookByIDInputDTO struct {
	ID string `validate:"required"`
}

type GetBookByIDOutputDTO struct {
	ID      string   `validate:"required"`
	Title   string   `validate:"required"`
	Authors []string `validate:"min=1"`
	Year    int      `validate:"gt=0"`
	Edition int      `validate:"gt=0"`
	Pages   int      `validate:"gt=0"`
}

type usecase struct {
	rep repository.BookRepository
}

func NewGetBookByID(rep repository.BookRepository) GetBookByID {
	return &usecase{rep}
}

func (u *usecase) Execute(input *GetBookByIDInputDTO) (*GetBookByIDOutputDTO, error) {

	check := validator.New()

	err := check.Struct(input)
	if err != nil {
		return nil, err
	}

	book, err := u.rep.SelectById(input.ID)

	if err != nil {
		return nil, errors.New("Error when getting book from repository")
	}

	output := &GetBookByIDOutputDTO{
		ID:      book.ID,
		Title:   book.Title,
		Authors: book.Authors,
		Year:    book.Year,
		Edition: book.Edition,
		Pages:   book.Pages,
	}

	err = check.Struct(output)
	if err != nil {
		return nil, err
	}

	return output, nil
}
