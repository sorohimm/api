package controllers

import (
	"api/internal/interfaces"
	"api/internal/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
)

type BookControllers struct {
	Log         *zap.SugaredLogger
	BookService interfaces.IBookService
}

func (c *BookControllers) CreateBook(ctx *gin.Context) {
	var book models.Book

	err := ctx.BindJSON(&book)
	if err != nil {
		c.Log.Warn(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"BindJson": err.Error()})
	}
	validate := validator.New()
	err = validate.Struct(&book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Validate": err.Error()})
	}

	//TODO: validate

	id, err := c.BookService.CreateBook(book)
	if err != nil {
		c.Log.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"CreateBook": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":   id,
		"book": book,
	})
}

func (c *BookControllers) GetBook(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "GetBook")
	}

	book, err := c.BookService.GetBook(id)
	if err != nil {
		c.Log.Error(err.Error())
		ctx.JSON(http.StatusNotFound, "GetBook")
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": book,
	})
}

func (c *BookControllers) GetAllBooks(ctx *gin.Context) {
	books, err := c.BookService.GetAllBooks()
	if err != nil {
		c.Log.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, "GetAllBooks")
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": books,
	})
}

func (c *BookControllers) DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.BookService.DeleteBook(id)
	if err != nil {
		c.Log.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, "DeleteBook")
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully deleted book: %s", id),
	})
}
