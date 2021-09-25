package interfaces

import (
	"api/internal/models"
)

type IBookService interface {
	CreateBook(models.Book) (models.BookResponse, error)
	GetBook(string) (models.BookResponse, error)
	GetAllBooks() ([]models.BookResponse, error)
	DeleteBook(string) error
}
