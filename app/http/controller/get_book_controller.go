package controller

import (
	"library/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetBookControllerOutput struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
	Year    int      `json:"year"`
	Edition int      `json:"edition"`
	Pages   int      `json:"pages"`
}

func GetBookController(u usecases.GetBookUseCase) gin.HandlerFunc {
	return func(g *gin.Context) {

		id := g.Param("id")

		UseCaseOutput, err := u.Execute(
			&usecases.GetBookUseCaseInputDTO{
				ID: id,
			})

		if err != nil {
			g.JSON(http.StatusBadRequest, outputMessage{err.Error()})
			return
		}

		if UseCaseOutput == nil {
			g.JSON(http.StatusNotFound, bookNotFound)
			return
		}

		output := GetBookControllerOutput{
			ID:      UseCaseOutput.ID,
			Title:   UseCaseOutput.Title,
			Authors: UseCaseOutput.Authors,
			Pages:   UseCaseOutput.Pages,
			Year:    UseCaseOutput.Year,
			Edition: UseCaseOutput.Edition,
		}

		g.JSON(http.StatusOK, output)
	}
}
