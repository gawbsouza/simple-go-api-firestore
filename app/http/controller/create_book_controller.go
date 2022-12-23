package controller

import (
	"library/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createBookControllerInput struct {
	Title   string   `json:"title" binding:"required"`
	Authors []string `json:"authors" binding:"required"`
	Year    int      `json:"year" binding:"required"`
	Edition int      `json:"edition" binding:"required"`
	Pages   int      `json:"pages" binding:"required"`
}

type createBookControllerOutput struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
	Year    int      `json:"year"`
	Edition int      `json:"edition"`
	Pages   int      `json:"pages"`
}

func CreateBookController(u usecases.CreateBookUseCase) gin.HandlerFunc {
	return func(g *gin.Context) {

		var input createBookControllerInput

		if err := g.ShouldBindJSON(&input); err != nil {
			g.JSON(http.StatusBadRequest, invalidJsonFormatError)
			return
		}

		UseCaseOutput, err := u.Execute(
			&usecases.CreateBookUseCaseInputDTO{
				Title:   input.Title,
				Authors: input.Authors,
				Year:    input.Year,
				Pages:   input.Pages,
				Edition: input.Edition,
			})

		if err != nil {
			g.JSON(http.StatusBadRequest, outputMessage{err.Error()})
			return
		}

		output := &createBookControllerOutput{
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
