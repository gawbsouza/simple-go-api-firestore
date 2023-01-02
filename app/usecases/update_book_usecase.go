package usecases

//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=usecases

import (
	"errors"
	"library/entity"

	"github.com/go-playground/validator/v10"
)

type UpdateBookUseCase interface {
	Execute(*UpdateBookUseCaseInputDTO) (*UpdateBookUseCaseOutputDTO, error)
}

type UpdateBookUseCaseInputDTO struct {
	ID      string   `validate:"required"`
	Title   string   `validate:"required"`
	Authors []string `validate:"min=1"`
	Year    int      `validate:"gt=0"`
	Edition int      `validate:"gt=0"`
	Pages   int      `validate:"gt=0"`
}

type UpdateBookUseCaseOutputDTO struct {
	ID      string
	Title   string
	Authors []string
	Year    int
	Edition int
	Pages   int
}

type UpdateBookUseCaseRepository interface {
	Update(*entity.Book) error
}

type updateBookUseCase struct {
	repository UpdateBookUseCaseRepository
}

func NewUpdateBookUseCase(rep UpdateBookUseCaseRepository) UpdateBookUseCase {
	return &updateBookUseCase{rep}
}

func (u *updateBookUseCase) Execute(input *UpdateBookUseCaseInputDTO) (*UpdateBookUseCaseOutputDTO, error) {

	validate := validator.New()
	err := validate.Struct(input)

	if err != nil {
		return nil, err
	}

	toUpdate := &entity.Book{
		ID:      input.ID,
		Title:   input.Title,
		Authors: input.Authors,
		Year:    input.Year,
		Edition: input.Edition,
		Pages:   input.Pages,
	}

	err = u.repository.Update(toUpdate)
	if err != nil {
		if err.Error() == "Book not found" {
			return nil, err
		}
		return nil, errors.New("Error when updating book")
	}

	output := &UpdateBookUseCaseOutputDTO{
		ID:      toUpdate.ID,
		Title:   toUpdate.Title,
		Authors: toUpdate.Authors,
		Year:    toUpdate.Year,
		Edition: toUpdate.Edition,
		Pages:   toUpdate.Pages,
	}

	return output, nil
}
