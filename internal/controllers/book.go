package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	"api/internal/services"
)

type BookControllers struct {
}

func CreateNewBook(c *gin.Context) {
	name, err := services.CreateNewBook(c)
	if err != nil {

	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf(" %s successfully created", name),
	})
}

func GetAllBooks(c *gin.Context) {
	books, err := services.GetAllBooks(c)
	if err != nil {

	}

	c.JSON(http.StatusOK, gin.H{
		"result": books,
	})
}

func GetBook(c *gin.Context) {
	book, err := services.GetBook(c)
	if err != nil {

	}

	c.JSON(http.StatusOK, gin.H{
		"result": book,
	})
}

func DeleteBook(c *gin.Context) {
	id, err := services.DeleteBook(c)
	if err != nil {

	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully deleted book: %s", id),
	})
}
