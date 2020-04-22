package routes

import (
	"github.com/go-chi/chi"
	book "github.com/mustanish/omelette/app/controllers"
	"github.com/mustanish/omelette/app/middlewares"
)

// Book is used to expose book specific routes
func Book() *chi.Mux {
	router := chi.NewRouter()
	router.With(middlewares.VerifyToken).With(middlewares.Validate).Post("/", book.AddBook)
	router.With(middlewares.VerifyToken).With(middlewares.Validate).Patch("/{id}", book.UpdateBook)
	router.With(middlewares.VerifyToken).Delete("/{id}", book.DeleteBook)
	router.With(middlewares.VerifyToken).Get("/{id}", book.SingleBook)
	router.With(middlewares.VerifyToken).Get("/all", book.AllBooks)
	return router
}
