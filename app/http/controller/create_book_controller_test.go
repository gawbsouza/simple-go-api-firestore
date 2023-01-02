package controller_test

import (
	"encoding/json"
	"library/http/controller"
	"library/usecases"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateBookController(t *testing.T) {

	t.Run("Invalid request body", func(t *testing.T) {

		ctlr := gomock.NewController(t)
		defer ctlr.Finish()

		mockUseCase := usecases.NewMockCreateBookUseCase(ctlr)
		mockUseCase.EXPECT().Execute(gomock.Any()).Times(0)

		g := gin.Default()
		g.POST("/books", controller.CreateBookController(mockUseCase))

		j, err := json.Marshal(&controller.CreateBookControllerInput{})

		if err != nil {
			t.Errorf("Error when creating input json")
		}

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/books", strings.NewReader(string(j)))

		g.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Use case error", func(t *testing.T) {

		ctlr := gomock.NewController(t)
		defer ctlr.Finish()

		mockUseCase := usecases.NewMockCreateBookUseCase(ctlr)
		mockUseCase.EXPECT().Execute(gomock.Any()).Times(0)

		g := gin.Default()
		g.POST("/books", controller.CreateBookController(mockUseCase))

		j, err := json.Marshal(&controller.CreateBookControllerInput{})

		if err != nil {
			t.Errorf("Error when creating input json")
		}

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/books", strings.NewReader(string(j)))

		g.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Success response body", func(t *testing.T) {

		ctlr := gomock.NewController(t)
		defer ctlr.Finish()

		mockUseCase := usecases.NewMockCreateBookUseCase(ctlr)
		mockUseCase.EXPECT().
			Execute(gomock.Eq(&usecases.CreateBookUseCaseInputDTO{
				Title:   "Book Title",
				Authors: []string{"Book Author"},
				Pages:   42,
				Year:    2022,
				Edition: 1,
			})).
			Return(&usecases.CreateBookUseCaseOutputDTO{
				ID:      "new-book-id",
				Title:   "Book Title",
				Authors: []string{"Book Author"},
				Pages:   42,
				Year:    2022,
				Edition: 1,
			}, nil).
			Times(1)

		g := gin.Default()
		g.POST("/books", controller.CreateBookController(mockUseCase))

		j, err := json.Marshal(&controller.CreateBookControllerInput{
			Title:   "Book Title",
			Authors: []string{"Book Author"},
			Pages:   42,
			Year:    2022,
			Edition: 1,
		})

		if err != nil {
			t.Errorf("Error when creating input json")
		}

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/books", strings.NewReader(string(j)))

		g.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var bodyJson controller.CreateBookControllerOutput
		err = json.NewDecoder(w.Body).Decode(&bodyJson)

		if err != nil {
			t.Errorf("Invalid output json format")
		}

		assert.Equal(t, "new-book-id", bodyJson.ID)
		assert.Equal(t, "Book Title", bodyJson.Title)
		assert.Equal(t, []string{"Book Author"}, bodyJson.Authors)
		assert.Equal(t, 2022, bodyJson.Year)
		assert.Equal(t, 1, bodyJson.Edition)
		assert.Equal(t, 42, bodyJson.Pages)
	})
}
