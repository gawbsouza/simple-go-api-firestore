package usecases_test

import (
	"errors"
	"library/entity"
	"library/usecases"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetBookUseCase_Execute(t *testing.T) {

	t.Run("Book not found", func(t *testing.T) {

		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepository := usecases.NewMockGetBookUseCaseRepository(controller)
		mockRepository.
			EXPECT().
			SelectById(gomock.Eq("not-found-id")).
			Return(nil, errors.New("Book not found"))

		uc := usecases.NewGetBookUseCase(mockRepository)
		res, err := uc.Execute(&usecases.GetBookUseCaseInputDTO{ID: "not-found-id"})

		assert.NotNil(t, err)
		assert.Equal(t, "Book not found", err.Error())
		assert.Nil(t, res)
	})

	t.Run("Book found", func(t *testing.T) {

		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepository := usecases.NewMockGetBookUseCaseRepository(controller)
		mockRepository.
			EXPECT().
			SelectById(gomock.Eq("found-id")).
			Return(&entity.Book{
				ID:      "found-id",
				Title:   "Book Title",
				Authors: []string{"Book Author"},
				Pages:   42,
				Edition: 1,
				Year:    2022,
			}, nil)

		uc := usecases.NewGetBookUseCase(mockRepository)
		res, err := uc.Execute(&usecases.GetBookUseCaseInputDTO{ID: "found-id"})

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, "found-id", res.ID)
		assert.Equal(t, "Book Title", res.Title)
		assert.Equal(t, []string{"Book Author"}, res.Authors)
		assert.Equal(t, 42, res.Pages)
		assert.Equal(t, 1, res.Edition)
		assert.Equal(t, 2022, res.Year)
	})

	t.Run("Repository Error", func(t *testing.T) {

		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepository := usecases.NewMockGetBookUseCaseRepository(controller)
		mockRepository.
			EXPECT().
			SelectById(gomock.Any()).
			Return(nil, errors.New("Repository Error"))

		uc := usecases.NewGetBookUseCase(mockRepository)
		res, err := uc.Execute(&usecases.GetBookUseCaseInputDTO{ID: "some-id"})

		assert.NotNil(t, err)
		assert.NotEqual(t, "Repository Error", err.Error())
		assert.Nil(t, res)
	})

	t.Run("Invalid input DTO", func(t *testing.T) {

		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepository := usecases.NewMockGetBookUseCaseRepository(controller)
		mockRepository.
			EXPECT().
			SelectById(gomock.Any()).
			Times(0)

		uc := usecases.NewGetBookUseCase(mockRepository)
		res, err := uc.Execute(&usecases.GetBookUseCaseInputDTO{})

		assert.NotNil(t, err)
		assert.Nil(t, res)
	})
}
