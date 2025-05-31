package main

import (
	"github.com/gin-gonic/gin"
)

// Book represents data about a recorded book.
type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

// books slice to seed recorded book data
var books = []Book{
	{ID: "1", Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", Price: 12.99},
	{ID: "2", Title: "1984", Author: "George Orwell", Price: 9.50},
	{ID: "3", Title: "To Kill a Mockingbird", Author: "Harper Lee", Price: 11.25},
}

func main() {
	router := gin.Default()

	// Define routes
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.POST("/books", addBook)
	router.PUT("/books/:id", updateBook)
	router.DELETE("/book/:id", deleteBook)

	router.Run("localhost:8080") // Listen and serve on 0.0.0.0:8080
}
