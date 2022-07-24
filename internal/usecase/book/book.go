package book

import (
	"awesomeProject1/internal/controller/http/dto"
	"awesomeProject1/internal/domain/entity"
	"context"
)

type BookService interface {
	GetAllForList(ctx context.Context) []*entity.BookView
	GetByID(ctx context.Context, id string) *entity.Book
}

type AuthorService interface {
	GetByID(ctx context.Context, id string) *entity.Author
}

type GenreService interface {
	GetByID(ctx context.Context, id string) *entity.Genre
}

type bookUsecase struct {
	bookService   BookService
	authorService AuthorService
	genreService  GenreService
}

func (u bookUsecase) CreateBook(ctx context.Context, dto dto.CreateBookDTO) (string, error) {
	return "", nil
}

func (u bookUsecase) ListAllBooks(ctx context.Context) []*entity.BookView {
	return u.bookService.GetAllForList(ctx)
}

func (u *bookUsecase) GetFullBook(ctx context.Context, id string) *entity.FullBook {
	book := u.bookService.GetByID(ctx, id)
	genre := u.genreService.GetByID(ctx, book.GenreID)
	author := u.authorService.GetByID(ctx, book.AuthorID)

	return entity.FullBook{
		Book:   book,
		Author: author,
		Genre:  genre,
	}
}
