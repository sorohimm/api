package interfaces

import (
	"api/internal/models"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type IDBBookRepo interface {
	InsertBook(context.Context, *pgxpool.Conn, models.Book) (models.Book, error)
	PullBook(context.Context, *pgxpool.Conn, string) (models.Book, error)
	PullAllBooks(context.Context, *pgxpool.Conn) ([]models.Book, error)
	DeleteBook(context.Context, *pgxpool.Conn, string) error
	UpdateBook(context.Context, *pgxpool.Conn, models.Book, string) (models.Book, error)
}
