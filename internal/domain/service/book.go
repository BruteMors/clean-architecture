package service

import (
	"awesomeProject1/internal/domain/entity"
	"context"
)

type bookStorage interface {
	GetOne(id string) *entity.Book
	GetAll(limit, offset int) []*entity.Book
	Create(book *entity.Book) *entity.Book
	Delete(book *entity.Book) error
}

type bookService struct {
	storage bookStorage
}

func NewBookService(storage bookStorage) *bookService {
	return &bookService{storage: storage}
}

func (s *bookService) Create(ctx context.Context) *entity.Book {
	return nil
}

func (s *bookService) GetByID(ctx context.Context, id string) *entity.Book {
	return s.storage.GetOne(id)
}

func (s *bookService) GetAll(ctx context.Context, limit, offset int) []*entity.Book {
	return s.storage.GetAll(limit, offset)
}

func (s *bookService) GetAllForList(ctx context.Context) []*entity.BookView {
	return nil
}
