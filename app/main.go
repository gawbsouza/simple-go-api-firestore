package main

import (
	"library/http/controller"
	repository "library/repository/firestore"
	"library/usecases"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	// Repositories
	bookRepository := repository.NewFireStoreBookRepository(os.Getenv("FIREBASE_PROJECT_ID"))

	// Usecases
	getAllBooksUseCase := usecases.NewGetAllBooksUseCase(bookRepository)
	getBookUseCase := usecases.NewGetBookUseCase(bookRepository)
	createBookUseCase := usecases.NewCreateBookUseCase(bookRepository)
	deleteBookUseCase := usecases.NewDeleteBookUseCase(bookRepository)
	updateBookUseCase := usecases.NewUpdateBookUseCase(bookRepository)

	// Controllers
	getAllBooksController := controller.GetAllBooksController(getAllBooksUseCase)
	getBookController := controller.GetBookController(getBookUseCase)
	createBookController := controller.CreateBookController(createBookUseCase)
	deleteBookController := controller.DeleteBookController(deleteBookUseCase)
	updateBookController := controller.UpdateBookController(updateBookUseCase)

	r := gin.Default()

	// Routes
	r.GET("/books", getAllBooksController)
	r.GET("/books/:id", getBookController)
	r.POST("/books", createBookController)
	r.PUT("/books/:id", updateBookController)
	r.DELETE("/books/:id", deleteBookController)

	r.Run(":4042")

}
