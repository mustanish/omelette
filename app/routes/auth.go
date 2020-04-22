package routes

import (
	"github.com/go-chi/chi"
	user "github.com/mustanish/omelette/app/controllers"
	"github.com/mustanish/omelette/app/middlewares"
)

// Auth is used to expose auth specific routes
func Auth() *chi.Mux {
	router := chi.NewRouter()
	router.With(middlewares.Validate).Post("/auth", user.Authenticate)
	router.With(middlewares.VerifyToken).With(middlewares.Validate).Post("/login", user.Login)
	router.With(middlewares.VerifyToken).Delete("/logout", user.Logout)
	router.With(middlewares.VerifyToken).Post("/refresh/token", user.RefreshToken)
	return router
}
