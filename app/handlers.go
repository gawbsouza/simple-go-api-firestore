package main

import (
	"context"
	"fmt"
	"library/book"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func getFireStoreClient(ctx context.Context) *firestore.Client {
	client, err := firestore.NewClient(ctx, os.Getenv("FIREBASE_PROJECT_ID"))
	if err != nil {
		log.Fatal(err.Error())
	}
	return client
}

func getAllBooks(c *gin.Context) {

	ctx := context.Background()

	client := getFireStoreClient(ctx)
	defer client.Close()

	docs, err := client.Collection("Books").Documents(ctx).GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Fatal(err.Error())
	}

	var lib = book.Library{}
	var b book.Book

	for _, doc := range docs {
		err = doc.DataTo(&b)
		if err == nil {
			b.Id = doc.Ref.ID
			lib.Books = append(lib.Books, b)
		}
	}

	c.JSON(http.StatusOK, lib)
}

func getBook(c *gin.Context) {

	ctx := context.Background()

	client := getFireStoreClient(ctx)
	defer client.Close()

	id := c.Param("id")
	doc, err := client.Collection("Books").Doc(id).Get(c.Request.Context())
	if err != nil {
		if status.Code(err) == codes.NotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var b book.Book
	doc.DataTo(&b)
	b.Id = id
	c.JSON(http.StatusOK, b)
}

func createBook(c *gin.Context) {

	var b book.Book
	if err := c.BindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	client := getFireStoreClient(ctx)
	defer client.Close()

	id := uuid.NewString()
	b.Id = id

	_, err := client.Collection("Books").Doc(id).Create(c.Request.Context(), b)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	location := fmt.Sprintf("/books/%s", id)
	c.Header("Location", location)
	c.JSON(http.StatusCreated, b)
}

func updateBook(c *gin.Context) {

	var b book.Book
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	client := getFireStoreClient(ctx)
	defer client.Close()

	id := c.Param("id")
	b.Id = id

	_, err := client.Collection("Books").Doc(id).Set(c.Request.Context(), b)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	location := fmt.Sprintf("/books/%s", id)
	c.Header("Location", location)
	c.JSON(http.StatusOK, b)
}

func deleteBook(c *gin.Context) {

	id := c.Param("id")

	ctx := context.Background()
	client := getFireStoreClient(ctx)
	defer client.Close()

	collection := client.Collection("Books")
	_, err := collection.Doc(id).Delete(c.Request.Context())
	if err != nil {
		if status.Code(err) == codes.NotFound {
			c.JSON(http.StatusNotModified, gin.H{"error": "Book not found."})
			return
		}
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}
