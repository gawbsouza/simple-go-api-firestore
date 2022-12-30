package usecases_test

import (
	"errors"
	"library/entity"
	"library/usecases"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateBookUseCase_Execute(t *testing.T) {

	t.Run("Success repository", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepository := usecases.NewMockCreateBookUseCaseRepository(controller)
		mockRepository.
			EXPECT().
			Insert(gomock.Any()).
			Return(&entity.Book{
				ID:      "Book ID",
				Title:   "Book Title",
				Authors: []string{"Book Author"},
				Year:    2022,
				Edition: 1,
				Pages:   42,
			}, nil)

		uc := usecases.NewCreateBookUseCase(mockRepository)
		res, err := uc.Execute(&usecases.CreateBookUseCaseInputDTO{
			Title:   "Book Title",
			Authors: []string{"Book Author"},
			Year:    2022,
			Edition: 1,
			Pages:   42,
		})

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.ID, "Book ID")
		assert.Equal(t, res.Title, "Book Title")
		assert.Equal(t, res.Authors, []string{"Book Author"})
		assert.Equal(t, res.Year, 2022)
		assert.Equal(t, res.Edition, 1)
		assert.Equal(t, res.Pages, 42)
	})

	t.Run("Error repository", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepository := usecases.NewMockCreateBookUseCaseRepository(controller)
		mockRepository.
			EXPECT().
			Insert(gomock.Any()).
			Return(nil, errors.New("Respository Error"))

		uc := usecases.NewCreateBookUseCase(mockRepository)
		res, err := uc.Execute(&usecases.CreateBookUseCaseInputDTO{
			Title:   "Book Title",
			Authors: []string{"Book Author"},
			Year:    2022,
			Edition: 1,
			Pages:   42,
		})

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "Error when inserting book into repository")
		assert.Nil(t, res)
	})

	t.Run("Invalid input DTO", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepository := usecases.NewMockCreateBookUseCaseRepository(controller)
		mockRepository.EXPECT().Insert(gomock.Any()).Times(0)

		uc := usecases.NewCreateBookUseCase(mockRepository)
		res, err := uc.Execute(&usecases.CreateBookUseCaseInputDTO{})

		assert.NotNil(t, err)
		assert.Nil(t, res)
	})
}
