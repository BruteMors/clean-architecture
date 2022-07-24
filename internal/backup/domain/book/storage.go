package book

import (
	"awesomeProject1/internal/domain/entity"
)

type Storage interface {
	GetOne(uuid string) *entity.Book
	GetAll(limit, offset int) []*entity.Book
	Create(book *entity.Book) *entity.Book
	Delete(book *entity.Book) error
}
