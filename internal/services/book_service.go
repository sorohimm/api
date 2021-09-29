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

func (s *BookService) CreateBook(book models.Book) (models.Book, error) {
	ctx := context.Background()
	conn, err := s.DBHandler.AcquireConn(ctx)
	if err != nil {
		return models.Book{}, err
	}
	defer conn.Release()

	result, err := s.DBBookRepo.InsertBook(ctx, conn, book)
	if err != nil {
		return models.Book{}, err
	}

	return result, nil
}

func (s *BookService) GetBook(id string) (models.Book, error) {
	ctx := context.Background()

	conn, err := s.DBHandler.AcquireConn(ctx)
	if err != nil {
		return models.Book{}, err
	}
	defer conn.Release()

	var result models.Book

	result, err = s.DBBookRepo.PullBook(ctx, conn, id)
	if err != nil {
		return models.Book{}, err
	}

	return result, nil

}

func (s *BookService) GetAllBooks() ([]models.Book, error) {
	ctx := context.Background()
	conn, err := s.DBHandler.AcquireConn(ctx)
	if err != nil {
		return nil, err
	}

	defer conn.Release()

	var books []models.Book

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

func (s *BookService) UpdateBook(book models.Book, uuid string) (models.Book, error) {
	ctx := context.Background()
	conn, err := s.DBHandler.AcquireConn(ctx)
	if err != nil {
		return models.Book{}, err
	}
	defer conn.Release()

	updResult, err := s.DBBookRepo.UpdateBook(ctx, conn, book, uuid)
	if err != nil {
		return models.Book{}, err
	}

	return updResult, nil
}