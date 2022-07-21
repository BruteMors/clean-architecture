package composites

import (
	"awesomeProject1/internal/adapters/api"
	"awesomeProject1/internal/adapters/api/author"
	author3 "awesomeProject1/internal/adapters/db/author"
	author2 "awesomeProject1/internal/domain/author"
)

type AuthorComposite struct {
	Storage author2.Storage
	Service author.Service
	Handler api.Handler
}

func NewAuthorComposite(mongoComposite *MongoDBComposite) (*AuthorComposite, error) {
	storage := author3.NewStorage(mongoComposite.db)
	service := author2.NewService(storage)
	handler := author.NewHandler(service)
	return &AuthorComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil
}
