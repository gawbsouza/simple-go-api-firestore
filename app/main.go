package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/books", getAllBooks)
	r.GET("/books/:id", getBook)
	r.POST("/books/", createBook)
	r.PUT("/books/:id", updateBook)
	r.DELETE("/books/:id", deleteBook)
	r.Run(":4042")
}
