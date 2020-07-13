package routes

import (
	"omelette/app/handlers"
	"omelette/app/middlewares"

	"github.com/go-chi/chi"
)

// Auth is used to expose auth specific routes
func Auth() *chi.Mux {
	router := chi.NewRouter()
	router.With(middlewares.Validate).Post("/auth", handlers.Authenticate)
	router.With(middlewares.VerifyToken).With(middlewares.Validate).Post("/login", handlers.Login)
	router.With(middlewares.VerifyToken).Delete("/logout", handlers.Logout)
	router.With(middlewares.VerifyToken).Post("/refresh/token", handlers.RefreshToken)
	return router
}
