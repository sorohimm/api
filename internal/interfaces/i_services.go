package interfaces

import (
	"api/internal/models"
)

type IBookService interface {
	CreateBook(models.Book) (string, error)
	GetBook(string) (models.Book, error)
	GetAllBooks() ([]models.Book, error)
	DeleteBook(string) (error)
}
