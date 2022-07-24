package service

import (
	author2 "awesomeProject1/internal/backup/domain/author"
	"awesomeProject1/internal/domain/entity"
	"context"
)

type AuthorStorage interface {
	GetOne(id string) *entity.Author
	GetAll(ctx context.Context) []*entity.Author
	Create(author *entity.Author) *entity.Author
	Delete(author *entity.Author) error
}

type authorService struct {
	storage AuthorStorage
}

func NewAuthorService(storage AuthorStorage) *authorService {
	return &authorService{storage: storage}
}

func (s *authorService) Create(ctx context.Context, dto *author2.CreateAuthorDTO) *entity.Author {
	return nil
}

func (s *authorService) GetByID(ctx context.Context, id string) *entity.Author {
	return s.storage.GetOne(id)
}

func (s *authorService) GetAll(ctx context.Context) []*entity.Author {
	return s.storage.GetAll(ctx)
}
