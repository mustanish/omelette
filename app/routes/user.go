package routes

import (
	"github.com/go-chi/chi"
	user "github.com/mustanish/omelette/app/controllers"
	"github.com/mustanish/omelette/app/middlewares"
)

// User is used to expose user specific routes
func User() *chi.Mux {
	r := chi.NewRouter()
	r.With(middlewares.ValidateBody).Post("/auth", user.Authenticate)
	r.With(middlewares.VerifyToken).With(middlewares.ValidateBody).Patch("/verify", user.VerifyUser)
	r.With(middlewares.VerifyToken).Patch("/", user.UpdateUser)
	r.With(middlewares.VerifyToken).Get("/", user.GetUser)
	r.With(middlewares.VerifyToken).Delete("/", user.Logout)
	return r
}
