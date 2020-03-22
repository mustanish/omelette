package routes

import (
	"github.com/go-chi/chi"
	book "github.com/mustanish/omelette/app/controllers"
)

// Book is used to expose book specific routes
func Book() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", book.AddBook)
	r.Get("/all-books", book.AllBooks)
	r.Route("/{id}", func(r chi.Router) {
		r.Patch("/{id}", book.UpdateBook)
		r.Get("/{id}", book.SingleBook)
		r.Delete("/{id}", book.DeleteBook)
	})
	return r
}
