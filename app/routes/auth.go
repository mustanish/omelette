package routes

import (
	"github.com/go-chi/chi"
	user "github.com/mustanish/omelette/app/controllers"
	"github.com/mustanish/omelette/app/middlewares"
)

// Auth is used to expose auth specific routes
func Auth() *chi.Mux {
	router := chi.NewRouter()
	router.With(middlewares.ValidateBody).Post("/auth", user.Authenticate)
	router.With(middlewares.VerifyToken).With(middlewares.ValidateBody).Post("/login", user.Login)
	router.With(middlewares.VerifyToken).Delete("/logout", user.Logout)
	router.With(middlewares.VerifyToken).Post("/refresh/token", user.RefreshToken)
	return router
}
