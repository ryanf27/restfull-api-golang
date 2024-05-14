package handlers

import (
	"net/http"

	"restfull-api/models"

	"github.com/gin-gonic/gin"

	"strconv"
)


func GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"books" :models.Books,
	})
}

func GetBookById(c *gin.Context) {
	idStr := c.Param("id")

	id , err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	book, found := findBookById(id)

	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"book": book})
	
}

func AddBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book data"})
		return
	}

	models.Books = append(models.Books, book)
	c.JSON(http.StatusCreated, gin.H{"message": "Book added successfully", "book": book})
}

func UpdateBook(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
        return
    }

    var updateBook models.Book
    if err := c.ShouldBindJSON(&updateBook); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book data"})
        return
    }

    index, found := findBookIndexById(id)
    if !found {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }

    models.Books[index] = updateBook
    c.JSON(http.StatusOK, gin.H{
        "message": "Book updated successfully",
        "book":    updateBook,
    })
}

func DeleteBook(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
        return
    }

    index, found := findBookIndexById(id)
    if !found {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }   

    models.Books = append(models.Books[:index], models.Books[index+1:]...)

    c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

func findBookById(id int) (models.Book, bool) {
    for _, book := range models.Books {
        bookID, err := strconv.Atoi(strconv.Itoa(book.ID))
        if err == nil && bookID == id {
            return book, true
        }
    }
    return models.Book{}, false
}

func findBookIndexById(id int) (int, bool) {
    for i, book := range models.Books {
        bookID, err := strconv.Atoi(strconv.Itoa(book.ID))
        if err == nil && bookID == id {
            return i, true
        }
    }
    return -1, false
}