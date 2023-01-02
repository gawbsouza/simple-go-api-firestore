package usecases_test

import (
	"errors"
	"library/usecases"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDeleteBookUseCase_Execute(t *testing.T) {

	t.Run("Invalid input dto", func(t *testing.T) {

		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepository := usecases.NewMockDeleteBookUseCaseRepository(controller)
		mockRepository.EXPECT().Delete(gomock.Any()).Times(0)

		uc := usecases.NewDeleteBookUseCase(mockRepository)
		err := uc.Execute(&usecases.DeleteBookUseCaseInputDTO{})

		assert.NotNil(t, err)
	})

	t.Run("Book not found", func(t *testing.T) {

		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepository := usecases.NewMockDeleteBookUseCaseRepository(controller)
		mockRepository.
			EXPECT().
			Delete(gomock.Eq("not-found-id")).
			Return(errors.New("Book not found"))

		uc := usecases.NewDeleteBookUseCase(mockRepository)
		err := uc.Execute(&usecases.DeleteBookUseCaseInputDTO{ID: "not-found-id"})

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "Book not found")
	})

	t.Run("Book found", func(t *testing.T) {

		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepository := usecases.NewMockDeleteBookUseCaseRepository(controller)
		mockRepository.
			EXPECT().
			Delete(gomock.Eq("found-id")).
			Return(nil)

		uc := usecases.NewDeleteBookUseCase(mockRepository)
		err := uc.Execute(&usecases.DeleteBookUseCaseInputDTO{ID: "found-id"})

		assert.Nil(t, err)
	})
}
