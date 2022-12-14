package usecases

//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=usecases

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type DeleteBookUseCase interface {
	Execute(*DeleteBookUseCaseInputDTO) error
}

type DeleteBookUseCaseInputDTO struct {
	ID string `validate:"required"`
}

type DeleteBookUseCaseRepository interface {
	Delete(id string) error
}

type deleteBookUseCase struct {
	repository DeleteBookUseCaseRepository
}

func NewDeleteBookUseCase(rep DeleteBookUseCaseRepository) DeleteBookUseCase {
	return &deleteBookUseCase{rep}
}

func (u *deleteBookUseCase) Execute(input *DeleteBookUseCaseInputDTO) error {

	validate := validator.New()

	err := validate.Struct(input)
	if err != nil {
		return err
	}

	err = u.repository.Delete(input.ID)
	if err != nil {
		if err.Error() == "Book not found" {
			return err
		}
		return errors.New("Error when deleting book from repository")
	}

	return err
}
