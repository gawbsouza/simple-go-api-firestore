package usecases

//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=usecases

import (
	"errors"
	"library/entity"

	"github.com/go-playground/validator/v10"
)

type GetBookUseCase interface {
	Execute(*GetBookUseCaseInputDTO) (*GetBookUseCaseOutputDTO, error)
}

type GetBookUseCaseInputDTO struct {
	ID string `validate:"required"`
}

type GetBookUseCaseOutputDTO struct {
	ID      string
	Title   string
	Authors []string
	Year    int
	Edition int
	Pages   int
}

type GetBookUseCaseRepository interface {
	SelectById(id string) (*entity.Book, error)
}

type getBookUseCase struct {
	repository GetBookUseCaseRepository
}

func NewGetBookUseCase(rep GetBookUseCaseRepository) GetBookUseCase {
	return &getBookUseCase{rep}
}

func (u *getBookUseCase) Execute(input *GetBookUseCaseInputDTO) (*GetBookUseCaseOutputDTO, error) {

	validate := validator.New()

	err := validate.Struct(input)
	if err != nil {
		return nil, err
	}

	book, err := u.repository.SelectById(input.ID)

	if err != nil {
		if err.Error() == "Book not found" {
			return nil, err
		}
		return nil, errors.New("Error when getting book from repository")
	}

	if book == nil {
		return nil, nil
	}

	output := &GetBookUseCaseOutputDTO{
		ID:      book.ID,
		Title:   book.Title,
		Authors: book.Authors,
		Year:    book.Year,
		Edition: book.Edition,
		Pages:   book.Pages,
	}

	return output, nil
}
