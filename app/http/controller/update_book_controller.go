package controller

import (
	"library/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type updateBookControllerInput struct {
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
	Year    int      `json:"year"`
	Edition int      `json:"edition"`
	Pages   int      `json:"pages"`
}

type updateBookControllerOutput struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
	Year    int      `json:"year"`
	Edition int      `json:"edition"`
	Pages   int      `json:"pages"`
}

func UpdateBookController(u usecases.UpdateBookUseCase) gin.HandlerFunc {
	return func(g *gin.Context) {

		id := g.Param("id")

		var input updateBookControllerInput

		if err := g.ShouldBindJSON(&input); err != nil {
			g.JSON(http.StatusBadRequest, invalidJsonFormatError)
			return
		}

		UseCaseOutput, err := u.Execute(
			&usecases.UpdateBookUseCaseInputDTO{
				ID:      id,
				Title:   input.Title,
				Authors: input.Authors,
				Year:    input.Year,
				Pages:   input.Pages,
				Edition: input.Edition,
			})

		if err != nil {
			if err.Error() == "Book not found" {
				g.JSON(http.StatusNotFound, bookNotFound)
				return
			}
			g.JSON(http.StatusBadRequest, outputMessage{err.Error()})
			return
		}

		output := &updateBookControllerOutput{
			ID:      UseCaseOutput.ID,
			Title:   UseCaseOutput.Title,
			Authors: UseCaseOutput.Authors,
			Year:    UseCaseOutput.Year,
			Pages:   UseCaseOutput.Pages,
			Edition: UseCaseOutput.Edition,
		}

		g.JSON(http.StatusOK, output)
	}
}
