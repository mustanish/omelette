package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/mustanish/omelette/app/responses"
)

var router = chi.NewRouter()

// InitializeRouter initializes router
func InitializeRouter() {
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // Set content-Type headers as application/json
		middleware.Logger,       // Log API request calls
		middleware.StripSlashes, // Strip slashes to no slash URL versions
		middleware.Recoverer,    // Recover from panics without crashing server
		cors.Handler,            // Enable CORS globally
	)
	router.Mount("/", Auth())
	router.Mount("/user", User())
	router.Mount("/book", Book())
	router.NotFound(responses.NotFound)
	router.MethodNotAllowed(responses.MethodNotAllowed)
}

// RouterInstance returns router instance
func RouterInstance() *chi.Mux {
	return router
}
