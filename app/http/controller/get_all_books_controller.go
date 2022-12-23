package controller

import (
	"library/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getAllBooksControllerOutput struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
	Year    int      `json:"year"`
	Edition int      `json:"edition"`
	Pages   int      `json:"pages"`
}

func GetAllBooksController(u usecases.GetAllBooksUseCase) gin.HandlerFunc {
	return func(g *gin.Context) {

		useCaseOutput, err := u.Execute()

		if err != nil {
			g.JSON(http.StatusBadRequest, internalServerError)
			return
		}

		if useCaseOutput == nil {
			g.JSON(http.StatusNotFound, outputMessage{"No books found"})
			return
		}

		var output []getAllBooksControllerOutput

		for _, book := range useCaseOutput {

			output = append(output, getAllBooksControllerOutput{
				ID:      book.ID,
				Title:   book.Title,
				Authors: book.Authors,
				Year:    book.Year,
				Pages:   book.Pages,
				Edition: book.Edition,
			})
		}

		g.JSON(http.StatusOK, output)
	}
}
