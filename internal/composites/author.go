package composites

import (
	author3 "awesomeProject1/internal/adapters/db/mongodb"
	"awesomeProject1/internal/backup/adapters/api"
	author2 "awesomeProject1/internal/backup/adapters/api/author"
	author4 "awesomeProject1/internal/backup/domain/author"
	"awesomeProject1/internal/controller/http/v1"
	service2 "awesomeProject1/internal/domain/service"
)

type AuthorComposite struct {
	Storage author4.Storage
	Service author2.Service
	Handler api.Handler
}

func NewAuthorComposite(mongoComposite *MongoDBComposite) (*AuthorComposite, error) {
	storage := author3.NewStorage(mongoComposite.db)
	service := service2.NewService(storage)
	handler := v1.NewHandler(service)
	return &AuthorComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil
}
