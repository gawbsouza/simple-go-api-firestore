package usecases

import (
	"errors"
	"library/entity"
)

type GetAllBooksUseCase interface {
	Execute() ([]*GetAllBooksUseCaseOutputDTO, error)
}

type GetAllBooksUseCaseOutputDTO struct {
	ID      string
	Title   string
	Authors []string
	Year    int
	Edition int
	Pages   int
}

type GetAllBooksUseCaseRepository interface {
	SelectAll() ([]entity.Book, error)
}

type getAllBooksUseCase struct {
	repository GetAllBooksUseCaseRepository
}

func NewGetAllBooksUseCase(rep GetAllBooksUseCaseRepository) GetAllBooksUseCase {
	return &getAllBooksUseCase{rep}
}

func (u *getAllBooksUseCase) Execute() ([]*GetAllBooksUseCaseOutputDTO, error) {

	books, err := u.repository.SelectAll()

	if err != nil {
		return nil, errors.New("Error when getting books from repository")
	}

	var output []*GetAllBooksUseCaseOutputDTO

	for _, book := range books {

		tmp := &GetAllBooksUseCaseOutputDTO{
			ID:      book.ID,
			Title:   book.Title,
			Authors: book.Authors,
			Year:    book.Year,
			Pages:   book.Pages,
			Edition: book.Edition,
		}

		output = append(output, tmp)
	}

	return output, nil
}
