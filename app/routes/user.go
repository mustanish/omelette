package routes

import (
	"github.com/go-chi/chi"
	user "github.com/mustanish/omelette/app/controllers"
	"github.com/mustanish/omelette/app/middlewares"
)

// User is used to expose user specific routes
func User() *chi.Mux {
	router := chi.NewRouter()
	router.With(middlewares.VerifyToken).With(middlewares.Validate).Patch("/", user.UpdateUser)
	router.With(middlewares.VerifyToken).Get("/", user.GetUser)
	return router
}
