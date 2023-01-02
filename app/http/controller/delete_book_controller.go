package controller

import (
	"library/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteBookController(u usecases.DeleteBookUseCase) gin.HandlerFunc {
	return func(g *gin.Context) {

		id := g.Param("id")

		err := u.Execute(&usecases.DeleteBookUseCaseInputDTO{ID: id})

		if err != nil {
			if err.Error() == "Book not found" {
				g.JSON(http.StatusNotFound, bookNotFound)
				return
			}
			g.JSON(http.StatusInternalServerError, internalServerError)
			return
		}

		g.JSON(http.StatusOK, successMessage)
	}
}
