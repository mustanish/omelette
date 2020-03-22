package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/mustanish/omelette/app/responses"
	"github.com/mustanish/omelette/app/routes"
)

func main() {
	r := chi.NewRouter()
	r.Use(
		render.SetContentType(render.ContentTypeJSON), // Set content-Type headers as application/json
		middleware.Logger,          // Log API request calls
		middleware.DefaultCompress, // Compress results, mostly gzipping assets and json
		middleware.RedirectSlashes, // Redirect slashes to no slash URL versions
		middleware.Recoverer,       // Recover from panics without crashing server
	)
	r.Mount("/user", routes.User())
	r.Mount("/book", routes.Book())
	r.NotFound(responses.NotFound)
	r.MethodNotAllowed(responses.MethodNotAllowed)
	fmt.Println("\033[32m" + "⇨ http server started at http://localhost:3333" + "\033[0m")
	http.ListenAndServe(":3333", r)
}
