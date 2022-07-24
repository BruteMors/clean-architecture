package composites

import (
	book3 "awesomeProject1/internal/adapters/db/mongodb"
	"awesomeProject1/internal/backup/adapters/api"
	book5 "awesomeProject1/internal/backup/adapters/api/book"
	book4 "awesomeProject1/internal/backup/domain/book"
	"awesomeProject1/internal/controller/http/v1"
	service2 "awesomeProject1/internal/domain/service"
)

type BookComposite struct {
	Storage book4.Storage
	Service book5.Service
	Handler api.Handler
}

func NewBookComposite(mongoComposite *MongoDBComposite) (*BookComposite, error) {
	storage := book3.NewStorage(mongoComposite.db)
	service := service2.NewService(storage)
	handler := v1.NewHandler(service)
	return &BookComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil
}
