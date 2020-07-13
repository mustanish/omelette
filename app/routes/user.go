package routes

import (
	"omelette/app/handlers"
	"omelette/app/middlewares"

	"github.com/go-chi/chi"
)

// User is used to expose user specific routes
func User() *chi.Mux {
	router := chi.NewRouter()
	router.With(middlewares.VerifyToken).With(middlewares.Validate).Patch("/", handlers.UpdateUser)
	router.With(middlewares.VerifyToken).Get("/", handlers.GetUser)
	return router
}
