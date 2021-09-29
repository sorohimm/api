package repos

import (
	"api/internal/models"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DBBookRepo struct {
}

func (r *DBBookRepo) InsertBook(
	ctx context.Context, conn *pgxpool.Conn, book models.Book) (models.Book, error) {
	const InsertBookStatement = `
	INSERT INTO books
	(name, year, author, category, price, descriptions)
	VALUES($1, $2, $3, $4, $5, $6)
	RETURNING "uuid";`

	err := conn.QueryRow(ctx, InsertBookStatement, book.Name, book.Year, book.Author,
		book.Category, book.Price, book.Descriptions).Scan(&book.Uuid)
	if err != nil {
		return models.Book{}, err
	}

	return book, err
}

func (r *DBBookRepo) PullBook(
	ctx context.Context, conn *pgxpool.Conn, uuid string) (book models.Book, err error) {
	const GetBookStatement = `SELECT * 
							  FROM books 
							  WHERE uuid = $1;`
	err = conn.QueryRow(ctx, GetBookStatement, uuid).Scan(&book.Uuid, &book.Name, &book.Year, &book.Author,
		&book.Category, &book.Price, &book.Descriptions)
	if err != nil {
		return models.Book{}, err
	}

	return book, err
}

func (r *DBBookRepo) PullAllBooks(ctx context.Context, conn *pgxpool.Conn) ([]models.Book, error) {
	const GetBooksStatement = `SELECT * 
							   FROM books;`
	rows, err := conn.Query(ctx, GetBooksStatement)

	if err != nil {
		return nil, err
	}

	var books []models.Book

	for rows.Next() {
		var book models.Book
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
	const DeleteBookStatement = `DELETE FROM books 
								 WHERE uuid = $1;`
	_, err := conn.Exec(ctx, DeleteBookStatement, uuid)
	if err != nil {
		return err
	}

	return nil
}

func (r *DBBookRepo) UpdateBook(ctx context.Context, conn *pgxpool.Conn, book models.Book, uuid string) (models.Book, error) {
	const UpdateBookStatement = `UPDATE books 
								 SET name = $1, year = $2, author = $3,category = $4, price = $5, descriptions = $6 
								 WHERE uuid = $7
								 RETURNING "uuid";`


	err := conn.QueryRow(ctx, UpdateBookStatement, book.Name, book.Year, book.Author,
		book.Category, book.Price, book.Descriptions, uuid).Scan(&book.Uuid)
	if err != nil {
		return models.Book{}, err
	}

	return book, err
}
