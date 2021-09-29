package interfaces

import (
	"api/internal/models"
)

type IBookService interface {
	CreateBook(models.Book) (models.Book, error)
	GetBook(string) (models.Book, error)
	GetAllBooks() ([]models.Book, error)
	DeleteBook(string) error
	UpdateBook(models.Book, string) (models.Book, error)
}
