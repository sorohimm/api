package repos

import (
	"api/internal/models"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DBBookRepo struct {
}

func (r *DBBookRepo) InsertBook(
	ctx context.Context, conn *pgxpool.Conn, book models.Book) (uuid string, err error) {
	const InsertBookStatement = `
	INSERT INTO books
	(name, year, author, category, price, descriptions)
	VALUES($1, $2, $3, $4, $5, $6)
	RETURNING "uuid";`
	err = conn.QueryRow(ctx, InsertBookStatement, book.Name, book.Year, book.Author,
		book.Category, book.Price, book.Descriptions).Scan(&uuid)

	return uuid, err
}

func (r *DBBookRepo) GetBook(
	ctx context.Context, conn *pgxpool.Conn, id string) (book models.Book, err error) {
	const GetBookStatement = `SELECT (id, name, year, author, category, price, descriptions) FROM books WHERE id = $1;`
	err = conn.QueryRow(ctx, GetBookStatement, id).Scan(&book.ID, &book.Name, &book.Year, &book.Author, &book.Category, &book.Price, &book.Descriptions)
	if err != nil {
		return models.Book{}, err
	}

	return book, err
}

func (r *DBBookRepo) GetAllBooks(ctx context.Context, conn *pgxpool.Conn) ([]models.Book, error) {
	//(id, name, year, author, category, price, descriptions)
	const GetBooksStatement = `SELECT * FROM books;`
	rows, err := conn.Query(ctx, GetBooksStatement)
	if err != nil {
		return nil, err
	}

	var books []models.Book

	for rows.Next() {
		var book models.Book
		err = rows.Scan(&book.ID, &book.Name, &book.Year, &book.Author, &book.Category, &book.Price, &book.Descriptions)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	defer rows.Close()

	return books, nil
}

func (r *DBBookRepo) DeleteBook(ctx context.Context, conn *pgxpool.Conn, id string) (err error) {
	const DeleteBookStatement = `DELETE FROM books WHERE id= $1;`
	_, err = conn.Exec(ctx, DeleteBookStatement, id)
	if err != nil {
		return err
	}

	return nil
}
