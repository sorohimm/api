package interfaces

import (
	"api/internal/models"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type IDBBookRepo interface {
	InsertBook(context.Context, *pgxpool.Conn, models.Book) (models.BookResponse, error)
	PullBook(context.Context, *pgxpool.Conn, string) (models.BookResponse, error)
	PullAllBooks(context.Context, *pgxpool.Conn) ([]models.BookResponse, error)
	DeleteBook(context.Context, *pgxpool.Conn, string) error
	UpdateBook(context.Context, *pgxpool.Conn, models.Book, string) (models.BookResponse, error)
}
