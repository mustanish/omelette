package responses

import (
	"net/http"
	"omelette/app/constants"

	"github.com/go-chi/render"
)

// HTTPSucess represents a success that occurred while handling a request.
type HTTPSucess struct {
	Code   int         `json:"-"` // http response status code
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

// Render is implemented for managing response payloads.
func (e *HTTPSucess) Render(res http.ResponseWriter, req *http.Request) error {
	render.Status(req, e.Code)
	return nil
}

// NewHTTPSucess sets the HTTPSucess struct while handling a request.
func NewHTTPSucess(code int, data interface{}) render.Renderer {
	return &HTTPSucess{Code: code, Status: constants.Success, Data: data}
}
