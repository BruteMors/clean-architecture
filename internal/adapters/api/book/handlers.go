package book

import (
	"awesomeProject1/internal/adapters/api"
	"context"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	bookURL  = "/books/:book_id"
	booksURL = "/books"
)

type handler struct {
	bookService Service
}

func NewHandler(service Service) api.Handler {
	return &handler{bookService: service}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(booksURL, h.GetAllBooks)
}

func (h *handler) GetAllBooks(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.bookService.GetAll(context.Background(), 0, 0)
	w.Write([]byte("books"))
	w.WriteHeader(http.StatusOK)
}