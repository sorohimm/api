package services

import (
	"api/internal/config"
	"api/internal/db"
	"api/internal/models"
	"context"
	"github.com/gin-gonic/gin"
)

func CreateNewBook(c *gin.Context) (string, error) {
	var newBook models.Book

	bindErr := c.BindJSON(&newBook)
	if bindErr != nil {
		return "", bindErr
	}

	cfg, cfgErr := config.New()
	if cfgErr != nil {
		return "", cfgErr
	}

	pool, err := db.InitDBClient(cfg, c)
	if err != nil {
		return "", err
	}

	_, createErr := pool.GetPool().Exec(context.Background(), "INSERT INTO book (name, year, author, category, price, descriptions)",
		newBook.Name, newBook.Year, newBook.Author, newBook.Category, newBook.Price, newBook.Descriptions)

	if createErr != nil {
		return newBook.Name, err
	}

	return newBook.Name, nil
}

func GetBook(c *gin.Context) (gin.H, error) {
	var (
		book   models.Book
		result gin.H
	)

	id := c.Param("id")

	cfg, cfgErr := config.New()
	if cfgErr != nil {
		return nil, cfgErr
	}

	pool, err := db.InitDBClient(cfg, c)
	if err != nil {
		return nil, err
	}

	err = pool.GetPool().QueryRow(context.Background(), "SELECT (id, name, year, author, category, price, descriptions) FROM book WHERE id = $1;", id).Scan(&book.ID)

	if err != nil {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
		return result, err
	}

	result = gin.H{
		"result": book,
	}
	return result, nil

}

func GetAllBooks(c *gin.Context) ([]models.Book, error) {
	var (
		book  models.Book
		books []models.Book
	)

	cfg, cfgErr := config.New()
	if cfgErr != nil {
		return nil, cfgErr
	}

	pool, err := db.InitDBClient(cfg, c)
	if err != nil {
		return nil, err
	}

	rows, selectErr := pool.GetPool().Query(context.Background(), "SELECT (id, name, year, author, category, price, descriptions) FROM book;")
	if selectErr != nil {
		return nil, selectErr
	}

	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Name, &book.Year, &book.Author, &book.Category, &book.Price, &book.Descriptions)
		books = append(books, book)
		if err != nil {
			return nil, err
		}
	}

	defer rows.Close()

	return books, nil
}

func DeleteBook(c *gin.Context) (string, error) {
	id := c.Query("id")

	cfg, cfgErr := config.New()
	if cfgErr != nil {
		return "", cfgErr
	}

	pool, err := db.InitDBClient(cfg, c)
	if err != nil {
		return "", err
	}

	_, rmErr := pool.GetPool().Exec(context.Background(), "DELETE FROM BOOK WHERE id= $1;", id)
	if rmErr != nil {
		return id, err
	}

	return id, nil
}
