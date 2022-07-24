package book

import (
	book2 "awesomeProject1/internal/backup/domain/book"
	"awesomeProject1/internal/domain/entity"
	"context"
)

type Service interface {
	GetByUUID(ctx context.Context, uuid string) *entity.Book
	GetAll(ctx context.Context, limit, offset int) []*entity.Book
	Create(ctx context.Context, dto *book2.CreateBookDTO) *entity.Book
}
