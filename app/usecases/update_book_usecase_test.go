package usecases_test

import (
	"errors"
	"library/usecases"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUpdateBookUseCase_Execute(t *testing.T) {

	t.Run("Book not found", func(t *testing.T) {

		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepository := usecases.NewMockUpdateBookUseCaseRepository(controller)
		mockRepository.
			EXPECT().
			Update(gomock.Any()).
			Return(errors.New("Book not found"))

		uc := usecases.NewUpdateBookUseCase(mockRepository)
		res, err := uc.Execute(&usecases.UpdateBookUseCaseInputDTO{
			ID:      "not-found-id",
			Title:   "Book Title",
			Authors: []string{"Book Author"},
			Pages:   42,
			Year:    2022,
			Edition: 1,
		})

		assert.NotNil(t, err)
		assert.Equal(t, "Book not found", err.Error())
		assert.Nil(t, res)
	})

	t.Run("Book found", func(t *testing.T) {

		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepository := usecases.NewMockUpdateBookUseCaseRepository(controller)
		mockRepository.
			EXPECT().
			Update(gomock.Any()).
			Return(nil)

		uc := usecases.NewUpdateBookUseCase(mockRepository)
		res, err := uc.Execute(&usecases.UpdateBookUseCaseInputDTO{
			ID:      "book-id",
			Title:   "Book Title",
			Authors: []string{"Book Author"},
			Pages:   42,
			Year:    2022,
			Edition: 1,
		})

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.ID, "book-id")
		assert.Equal(t, res.Title, "Book Title")
		assert.Equal(t, res.Authors, []string{"Book Author"})
		assert.Equal(t, res.Year, 2022)
		assert.Equal(t, res.Edition, 1)
		assert.Equal(t, res.Pages, 42)
	})

	t.Run("Repository Error", func(t *testing.T) {

		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepository := usecases.NewMockUpdateBookUseCaseRepository(controller)
		mockRepository.
			EXPECT().
			Update(gomock.Any()).
			Return(errors.New("Repository Error"))

		uc := usecases.NewUpdateBookUseCase(mockRepository)
		res, err := uc.Execute(&usecases.UpdateBookUseCaseInputDTO{
			ID:      "not-found-id",
			Title:   "Book Title",
			Authors: []string{"Book Author"},
			Pages:   42,
			Year:    2022,
			Edition: 1,
		})

		assert.NotNil(t, err)
		assert.NotEqual(t, "Repository Error", err.Error())
		assert.Nil(t, res)
	})

	t.Run("Invalid input DTO", func(t *testing.T) {

		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepository := usecases.NewMockUpdateBookUseCaseRepository(controller)
		mockRepository.
			EXPECT().
			Update(gomock.Any()).
			Times(0)

		uc := usecases.NewUpdateBookUseCase(mockRepository)
		res, err := uc.Execute(&usecases.UpdateBookUseCaseInputDTO{})

		assert.NotNil(t, err)
		assert.Nil(t, res)
	})
}
