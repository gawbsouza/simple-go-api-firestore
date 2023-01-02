package controller_test

import (
	"encoding/json"
	"library/http/controller"
	"library/usecases"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetBookController(t *testing.T) {

	t.Run("Book not found", func(t *testing.T) {

		ctlr := gomock.NewController(t)
		defer ctlr.Finish()

		mockUseCase := usecases.NewMockGetBookUseCase(ctlr)
		mockUseCase.
			EXPECT().
			Execute(gomock.Eq(&usecases.GetBookUseCaseInputDTO{ID: "invalid-book-id"})).
			Return(nil, nil)

		g := gin.Default()
		g.GET("/books/:id", controller.GetBookController(mockUseCase))

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/books/invalid-book-id", nil)

		g.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("Book found", func(t *testing.T) {

		ctlr := gomock.NewController(t)
		defer ctlr.Finish()

		mockUseCase := usecases.NewMockGetBookUseCase(ctlr)
		mockUseCase.
			EXPECT().
			Execute(gomock.Eq(&usecases.GetBookUseCaseInputDTO{ID: "valid-book-id"})).
			Return(&usecases.GetBookUseCaseOutputDTO{
				ID:      "valid-book-id",
				Title:   "Book Title",
				Authors: []string{"Book Author"},
				Pages:   42,
				Year:    2022,
				Edition: 1,
			}, nil)

		g := gin.Default()
		g.GET("/books/:id", controller.GetBookController(mockUseCase))

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/books/valid-book-id", nil)

		g.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var bodyJson controller.CreateBookControllerOutput
		err := json.NewDecoder(w.Body).Decode(&bodyJson)

		if err != nil {
			t.Errorf("Invalid output json format")
		}

		assert.Equal(t, "valid-book-id", bodyJson.ID)
		assert.Equal(t, "Book Title", bodyJson.Title)
		assert.Equal(t, []string{"Book Author"}, bodyJson.Authors)
		assert.Equal(t, 2022, bodyJson.Year)
		assert.Equal(t, 1, bodyJson.Edition)
		assert.Equal(t, 42, bodyJson.Pages)
	})
}
