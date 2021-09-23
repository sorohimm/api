package services

import (
	_"api/internal/config"
	_"api/internal/db"
	"api/internal/interfaces"
	"api/internal/models"
	"context"
	_"github.com/gin-gonic/gin"
)

type BookService struct {
	DBHandler interfaces.IDBHandler
	DBBookRepo interfaces.IDBBookRepo
}

func (s *BookService) CreateBook(book models.Book) (string, error) {
	ctx := context.Background()
	conn, err := s.DBHandler.AcquireConn(ctx)
	if err != nil {
		return "", err
	}

	id, err := s.DBBookRepo.InsertBook(ctx, conn, book)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *BookService) GetBook(id string) (models.Book, error) {
	ctx := context.Background()
	conn, err := s.DBHandler.AcquireConn(ctx)
	if err != nil {
		return models.Book{}, err
	}
	var result models.Book

	result, err = s.DBBookRepo.GetBook(ctx, conn, id)
	if err != nil {
		return models.Book{}, err
	}

	return result, nil

}

func (s *BookService) GetAllBooks() ([]models.Book, error) {
	ctx := context.Background()
	conn, err := s.DBHandler.AcquireConn(ctx)
	if err != nil  {
		return nil, err
	}

	var books []models.Book

	books, err = s.DBBookRepo.GetAllBooks(ctx, conn)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (s *BookService) DeleteBook(id string) (error) {
	ctx := context.Background()
	conn, err := s.DBHandler.AcquireConn(ctx)
	if err != nil  {
		return err
	}

	err = s.DBBookRepo.DeleteBook(ctx, conn, id)
	if err != nil  {
		return err
	}

	return nil
}
