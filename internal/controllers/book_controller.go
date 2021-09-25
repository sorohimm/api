package controllers

import (
	"api/internal/interfaces"
	"api/internal/models"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type BookControllers struct {
	Log         *zap.SugaredLogger
	BookService interfaces.IBookService
}

func (c *BookControllers) CreateBook(ctx *gin.Context) {
	var book models.Book

	validate := validator.New()
	err := validate.Struct(&book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Validate": "invalid request"})
		return
	}

	err = ctx.BindJSON(&book)
	if err != nil {
		c.Log.Warn(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"BindJson": "invalid json"})
		return
	}

	result, err := c.BookService.CreateBook(book)
	if err != nil {
		c.Log.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"CreateBook": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *BookControllers) GetBook(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := uuid.Parse(id)

	if err != nil {
		c.Log.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Validate": err.Error()})
		return
	}

	book, err := c.BookService.GetBook(id)
	if err != nil {
		c.Log.Error(err.Error())
		ctx.JSON(http.StatusNotFound, gin.H{
			"GetBook": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (c *BookControllers) GetAllBooks(ctx *gin.Context) {
	books, err := c.BookService.GetAllBooks()
	if err != nil {
		c.Log.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"PullAllBooks": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func (c *BookControllers) DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := uuid.Parse(id)

	if err != nil {
		c.Log.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Validate": err.Error()})
		return
	}

	err = c.BookService.DeleteBook(id)
	if err != nil {
		c.Log.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"DeleteBook": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully deleted book: %s", id),
	})
}
