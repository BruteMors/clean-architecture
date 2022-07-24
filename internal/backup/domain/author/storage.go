package author

import (
	"awesomeProject1/internal/domain/entity"
)

type Storage interface {
	GetOne(uuid string) *entity.Author
	GetAll(limit, offset int) []*entity.Author
	Create(author *entity.Author) *entity.Author
	Delete(author *entity.Author) error
}
