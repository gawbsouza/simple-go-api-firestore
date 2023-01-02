package controller_test

import (
	"encoding/json"
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

func TestGetAllBoksController(t *testing.T) {

	t.Run("No books", func(t *testing.T) {

		ctlr := gomock.NewController(t)
		defer ctlr.Finish()

		mockUseCase := usecases.NewMockGetAllBooksUseCase(ctlr)
		mockUseCase.EXPECT().Execute().Return(nil, nil)

		g := gin.Default()
		g.GET("/books", controller.GetAllBooksController(mockUseCase))

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/books", nil)

		g.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("Usecase error", func(t *testing.T) {

		ctlr := gomock.NewController(t)
		defer ctlr.Finish()

		mockUseCase := usecases.NewMockGetAllBooksUseCase(ctlr)
		mockUseCase.EXPECT().Execute().Return(nil, errors.New("Usecase Error"))

		g := gin.Default()
		g.GET("/books", controller.GetAllBooksController(mockUseCase))

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/books", nil)

		g.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("Return some books", func(t *testing.T) {

		ctlr := gomock.NewController(t)
		defer ctlr.Finish()

		mockUseCase := usecases.NewMockGetAllBooksUseCase(ctlr)
		mockUseCase.EXPECT().Execute().Return([]*usecases.GetAllBooksUseCaseOutputDTO{
			{
				ID:      "book-id-1",
				Title:   "Book Title",
				Authors: []string{"Book Author"},
				Pages:   42,
				Year:    2022,
				Edition: 1,
			},
			{
				ID:      "book-id-2",
				Title:   "Book Title",
				Authors: []string{"Book Author"},
				Pages:   42,
				Year:    2022,
				Edition: 1,
			},
		}, nil)

		g := gin.Default()
		g.GET("/books", controller.GetAllBooksController(mockUseCase))

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/books", nil)

		g.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var bodyJson []controller.GetAllBooksControllerOutput
		err := json.NewDecoder(w.Body).Decode(&bodyJson)

		if err != nil {
			t.Errorf("Invalid output json format")
		}

		assert.Equal(t, 2, len(bodyJson))

		assert.Equal(t, "book-id-1", bodyJson[0].ID)
		assert.Equal(t, "Book Title", bodyJson[0].Title)
		assert.Equal(t, []string{"Book Author"}, bodyJson[0].Authors)
		assert.Equal(t, 2022, bodyJson[0].Year)
		assert.Equal(t, 1, bodyJson[0].Edition)
		assert.Equal(t, 42, bodyJson[0].Pages)

		assert.Equal(t, "book-id-2", bodyJson[1].ID)
		assert.Equal(t, "Book Title", bodyJson[1].Title)
		assert.Equal(t, []string{"Book Author"}, bodyJson[1].Authors)
		assert.Equal(t, 2022, bodyJson[1].Year)
		assert.Equal(t, 1, bodyJson[1].Edition)
		assert.Equal(t, 42, bodyJson[1].Pages)
	})
}
