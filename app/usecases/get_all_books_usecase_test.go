package usecases_test

import (
	"errors"
	"library/entity"
	"library/usecases"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetAllBooksUsecase_Execute(t *testing.T) {

	t.Run("No books found", func(t *testing.T) {

		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepository := usecases.NewMockGetAllBooksUseCaseRepository(controller)
		mockRepository.EXPECT().SelectAll().Return(nil, nil)

		uc := usecases.NewGetAllBooksUseCase(mockRepository)
		res, err := uc.Execute()

		assert.Nil(t, err)
		assert.Nil(t, res)
	})

	t.Run("Some Books found", func(t *testing.T) {

		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepository := usecases.NewMockGetAllBooksUseCaseRepository(controller)
		mockRepository.
			EXPECT().
			SelectAll().
			Return([]entity.Book{{}, {}}, nil)

		uc := usecases.NewGetAllBooksUseCase(mockRepository)
		res, err := uc.Execute()

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, 2, len(res))
	})

	t.Run("Repository error", func(t *testing.T) {

		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepository := usecases.NewMockGetAllBooksUseCaseRepository(controller)
		mockRepository.
			EXPECT().
			SelectAll().
			Return(nil, errors.New("Repository Error"))

		uc := usecases.NewGetAllBooksUseCase(mockRepository)
		res, err := uc.Execute()

		assert.NotNil(t, err)
		assert.NotEqual(t, "Repository Error", err.Error())
		assert.Nil(t, res)
	})
}
