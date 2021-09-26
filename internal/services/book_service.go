package services

import (
	"api/internal/interfaces"
	"api/internal/models"
	"context"
)

type BookService struct {
	DBHandler  interfaces.IDBHandler
	DBBookRepo interfaces.IDBBookRepo
}

func (s *BookService) CreateBook(book models.Book) (models.BookResponse, error) {
	ctx := context.Background()
	conn, err := s.DBHandler.AcquireConn(ctx)
	if err != nil {
		return models.BookResponse{}, err
	}
	defer conn.Release()

	result, err := s.DBBookRepo.InsertBook(ctx, conn, book)
	if err != nil {
		return models.BookResponse{}, err
	}

	return result, nil
}

func (s *BookService) GetBook(id string) (models.BookResponse, error) {
	ctx := context.Background()

	conn, err := s.DBHandler.AcquireConn(ctx)
	if err != nil {
		return models.BookResponse{}, err
	}
	defer conn.Release()

	var result models.BookResponse

	result, err = s.DBBookRepo.PullBook(ctx, conn, id)
	if err != nil {
		return models.BookResponse{}, err
	}

	return result, nil

}

func (s *BookService) GetAllBooks() ([]models.BookResponse, error) {
	ctx := context.Background()
	conn, err := s.DBHandler.AcquireConn(ctx)
	if err != nil {
		return nil, err
	}

	defer conn.Release()

	var books []models.BookResponse

	books, err = s.DBBookRepo.PullAllBooks(ctx, conn)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (s *BookService) DeleteBook(id string) error {
	ctx := context.Background()
	conn, err := s.DBHandler.AcquireConn(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	err = s.DBBookRepo.DeleteBook(ctx, conn, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *BookService) UpdateBook(book models.Book, uuid string) (models.BookResponse, error) {
	ctx := context.Background()
	conn, err := s.DBHandler.AcquireConn(ctx)
	if err != nil {
		return models.BookResponse{}, err
	}
	defer conn.Release()

	respBook, err := s.GetBook(uuid)
	if err != nil {
		return models.BookResponse{}, err
	}

	var bookCpy = book

	if bookCpy.Name == "" {
		bookCpy.Name = respBook.Name
	}
	if bookCpy.Year == "" {
		bookCpy.Year = respBook.Year
	}
	if bookCpy.Author == "" {
		bookCpy.Author = respBook.Author
	}
	if bookCpy.Price == "" {
		bookCpy.Price = respBook.Price
	}
	if bookCpy.Descriptions == "" {
		bookCpy.Descriptions = respBook.Descriptions
	}

	result, err := s.DBBookRepo.UpdateBook(ctx, conn, bookCpy, uuid)
	if err != nil {
		return models.BookResponse{}, err
	}

	return result, nil
}