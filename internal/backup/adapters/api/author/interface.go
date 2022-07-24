package author

import (
	author2 "awesomeProject1/internal/backup/domain/author"
	"awesomeProject1/internal/domain/entity"
	"context"
)

type Service interface {
	GetByUUID(ctx context.Context, uuid string) *entity.Author
	GetAll(ctx context.Context, limit, offset int) []*entity.Author
	Create(ctx context.Context, dto *author2.CreateAuthorDTO) *entity.Author
}
