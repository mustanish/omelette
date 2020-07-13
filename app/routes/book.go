package routes

import (
	"omelette/app/handlers"
	"omelette/app/middlewares"

	"github.com/go-chi/chi"
)

// Book is used to expose book specific routes
func Book() *chi.Mux {
	router := chi.NewRouter()
	router.With(middlewares.VerifyToken).With(middlewares.Validate).Post("/", handlers.AddBook)
	router.With(middlewares.VerifyToken).With(middlewares.Validate).Patch("/{id}", handlers.UpdateBook)
	router.With(middlewares.VerifyToken).Delete("/{id}", handlers.DeleteBook)
	router.With(middlewares.VerifyToken).Get("/{id}", handlers.SingleBook)
	router.With(middlewares.VerifyToken).Get("/all", handlers.AllBooks)
	return router
}
