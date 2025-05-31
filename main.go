package main

import (
	"net/http"

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
	router.DELETE("/books/:id", deleteBook)

	router.Run("localhost:8080") // Listen and serve on 0.0.0.0:8080
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func getBookByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range books {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "book not found"})
}

func addBook(c *gin.Context) {
	var newBook Book

	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func updateBook(c *gin.Context) {
	id := c.Param("id")

	var updatedBook Book

	if err := c.BindJSON(&updatedBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	found := false
	for i, b := range books {
		if b.ID == id {
			books[i] = updatedBook
			books[i].ID = id
			c.IndentedJSON(http.StatusOK, books[i])
			found = true
			break
		}
	}

	if !found {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
	}
}

func deleteBook(c *gin.Context) {
	id := c.Param("id")

	found := false
	for i, b := range books {
		if b.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "book deleted"})
			found = true
			break
		}
	}

	if !found {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
	}
}
