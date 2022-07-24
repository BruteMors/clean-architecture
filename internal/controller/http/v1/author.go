package v1

import (
	"awesomeProject1/internal/backup/adapters/api"
	"awesomeProject1/internal/backup/adapters/api/author"
	"github.com/julienschmidt/httprouter"
)

type handler struct {
	authorService author.Service
}

func (h *handler) Register(router *httprouter.Router) {
	//TODO implement me
	panic("implement me")
}

func NewHandler(service author.Service) api.Handler {
	return &handler{authorService: service}
}
