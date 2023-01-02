package controller_test

import (
	"errors"
	"library/http/controller"
	"library/usecases"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDeleteBookController(t *testing.T) {

	t.Run("Invalid book ID", func(t *testing.T) {

		ctlr := gomock.NewController(t)
		defer ctlr.Finish()

		mockUseCase := usecases.NewMockDeleteBookUseCase(ctlr)
		mockUseCase.
			EXPECT().
			Execute(gomock.Eq(&usecases.DeleteBookUseCaseInputDTO{ID: "invalid-book-id"})).
			Return(errors.New("Book not found"))

		g := gin.Default()
		g.DELETE("/books/:id", controller.DeleteBookController(mockUseCase))

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/books/invalid-book-id", nil)

		g.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("Valid book ID", func(t *testing.T) {

		ctlr := gomock.NewController(t)
		defer ctlr.Finish()

		mockUseCase := usecases.NewMockDeleteBookUseCase(ctlr)
		mockUseCase.
			EXPECT().
			Execute(gomock.Eq(&usecases.DeleteBookUseCaseInputDTO{ID: "valid-book-id"})).
			Return(nil)

		g := gin.Default()
		g.DELETE("/books/:id", controller.DeleteBookController(mockUseCase))

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/books/valid-book-id", nil)

		g.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Usecase error", func(t *testing.T) {

		ctlr := gomock.NewController(t)
		defer ctlr.Finish()

		mockUseCase := usecases.NewMockDeleteBookUseCase(ctlr)
		mockUseCase.
			EXPECT().
			Execute(gomock.Any()).
			Return(errors.New("Usecase error"))

		g := gin.Default()
		g.DELETE("/books/:id", controller.DeleteBookController(mockUseCase))

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/books/some-book-id", nil)

		g.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
