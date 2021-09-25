package repos

import (
	"api/internal/models"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DBBookRepo struct {
}

func (r *DBBookRepo) InsertBook(
	ctx context.Context, conn *pgxpool.Conn, reqBook models.Book) (models.BookResponse, error) {
	const InsertBookStatement = `
	INSERT INTO books
	(name, year, author, category, price, descriptions)
	VALUES($1, $2, $3, $4, $5, $6)
	RETURNING "uuid";`

	respBook := models.BookResponse{
		Name:         reqBook.Uuid,
		Year:         reqBook.Year,
		Author:       reqBook.Author,
		Category:     reqBook.Category,
		Price:        reqBook.Price,
		Descriptions: reqBook.Descriptions,
	}

	err := conn.QueryRow(ctx, InsertBookStatement, reqBook.Uuid, reqBook.Year, reqBook.Author,
		reqBook.Category, reqBook.Price, reqBook.Descriptions).Scan(&respBook.Uuid)

	return respBook, err
}

func (r *DBBookRepo) PullBook(
	ctx context.Context, conn *pgxpool.Conn, uuid string) (book models.BookResponse, err error) {
	const GetBookStatement = `SELECT * FROM books WHERE uuid =$1;`
	err = conn.QueryRow(ctx, GetBookStatement, uuid).Scan(&book.Uuid, &book.Name, &book.Year, &book.Author,
		&book.Category, &book.Price, &book.Descriptions)
	if err != nil {
		return models.BookResponse{}, err
	}

	return book, err
}

func (r *DBBookRepo) PullAllBooks(ctx context.Context, conn *pgxpool.Conn) ([]models.BookResponse, error) {
	const GetBooksStatement = `SELECT * FROM books;`
	rows, err := conn.Query(ctx, GetBooksStatement)

	if err != nil {
		return nil, err
	}

	var books []models.BookResponse

	for rows.Next() {
		var book models.BookResponse
		err = rows.Scan(&book.Uuid, &book.Name, &book.Year, &book.Author, &book.Category, &book.Price, &book.Descriptions)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	defer rows.Close()

	return books, nil
}

func (r *DBBookRepo) DeleteBook(ctx context.Context, conn *pgxpool.Conn, uuid string) error {
	const DeleteBookStatement = `DELETE FROM books WHERE uuid=$1;`
	_, err := conn.Exec(ctx, DeleteBookStatement, uuid)
	if err != nil {
		return err
	}

	return nil
}
