package routes

import (
	"github.com/go-chi/chi"
	user "github.com/mustanish/omelette/app/controllers"
	"github.com/mustanish/omelette/app/middlewares"
)

// User is used to expose user specific routes
func User() *chi.Mux {
	r := chi.NewRouter()
	r.With(middlewares.Body).Post("/auth", user.Authenticate)
	r.Post("/verify", user.VerifyUser)
	r.Route("/{id}", func(r chi.Router) {
		r.Patch("/{id}", user.UpdateUser)
		r.Get("/{id}", user.GetUser)
		r.Delete("/{id}", user.DeleteUser)
	})
	return r
}
