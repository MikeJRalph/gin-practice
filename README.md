# gin-practice
Basic practice using Gin web framework.

## Project Idea: Simple Bookstore API

This project will allow you to manage a list of books with basic CRUD (Create, Read, Update, Delete) operations.

Core Features to Implement:
1. Get All Books: Retrieve a list of all books.
2. Get Single Book: Retrieve details of a specific book by its ID.
3. Add New Book: Create a new book entry.
4. Update Book: Modify an existing book's details.
5. Delete Book: Remove a book from the list.

### Data Structure (Go Struct):
You'll need a struct to represent a book.

### In-Memory Data Store:
For simplicity, use a slice of Book structs to store your data. This data will be lost when the application restarts, but it's perfect for practice.

Feature	    HTTP Method	Endpoint	Description
Get All Books	 GET	/books	    Returns a JSON array of all books.
Get Single Book	 GET	/books/:id	Returns a JSON object of a single book by ID.
Add New Book	POST	/books	    Creates a new book from JSON request body.
Update Book	     PUT	/books/:id	Updates an existing book by ID.
Delete Book	    DELETE	/books/:id	Deletes a book by ID.

## How to use

Launch the app:

```
go run .
```

Open a Git Bash terminal

Then enter the following:

Get all books:

```
curl http://localhost:8080/books
```

Get a single book:

```
curl http://localhost:8080/books/1
```

Add a new book:

```
curl -X POST -H "Content-Type: application/json" -d '{"id": "4", "title": "The Lord of the Rings", "author": "J.R.R. Tolkien", "price": 25.00}' http://localhost:8080/books
```

Update a book:

```
curl -v -X PUT -H "Content-Type: application/json" -d '{
    "id": "1",
    "title": "The Updated Book Title",
    "author": "New Author Name",
    "price": 19.99
}' http://localhost:8080/books/1
```

Delete a book:

```
curl -X DELETE http://localhost:8080/books/2
```