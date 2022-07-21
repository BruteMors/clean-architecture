package author

import (
	"awesomeProject1/internal/adapters/api"
	"github.com/julienschmidt/httprouter"
)

type handler struct {
	authorService Service
}

func (h *handler) Register(router *httprouter.Router) {
	//TODO implement me
	panic("implement me")
}

func NewHandler(service Service) api.Handler {
	return &handler{authorService: service}
}
