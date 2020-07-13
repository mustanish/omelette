package responses

import (
	"net/http"
	"omelette/app/constants"

	"github.com/go-chi/render"
)

// HTTPError represents an error that occurred while handling a request.
type HTTPError struct {
	Code   int         `json:"-"` // http response status code
	Status string      `json:"status"`
	Error  interface{} `json:"error"`
}

// Render is implemented for managing response payloads.
func (e *HTTPError) Render(res http.ResponseWriter, req *http.Request) error {
	render.Status(req, e.Code)
	return nil
}

// NotFound custom handler for endpoints that could not be found.
func NotFound(res http.ResponseWriter, req *http.Request) {
	render.Render(res, req, NewHTTPError(http.StatusNotFound, constants.NotFound))
}

// MethodNotAllowed custom handler for endpoints where the method is unresolved.
func MethodNotAllowed(res http.ResponseWriter, req *http.Request) {
	render.Render(res, req, NewHTTPError(http.StatusMethodNotAllowed, constants.MethodNotAllowed))
}

// NewHTTPError sets the HTTPError struct while handling a request.
func NewHTTPError(code int, error interface{}) render.Renderer {
	return &HTTPError{Code: code, Status: constants.Failed, Error: error}
}
