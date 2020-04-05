package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/mustanish/omelette/app/config"
	"github.com/mustanish/omelette/app/responses"
	"github.com/mustanish/omelette/app/routes"
)

func main() {
	var (
		router    = chi.NewRouter()
		config, _ = config.LoadConfig()
	)
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
	router.Mount("/", routes.Auth())
	router.Mount("/user", routes.User())
	router.Mount("/book", routes.Book())
	router.NotFound(responses.NotFound)
	router.MethodNotAllowed(responses.MethodNotAllowed)
	fmt.Println("\033[32m" + "⇨ http server started at " + config.Server.Host + ":" + config.Server.Port + "\033[0m")
	http.ListenAndServe(":"+config.Server.Port, router)
}
