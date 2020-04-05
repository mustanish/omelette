package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-chi/render"
	"github.com/mustanish/omelette/app/constants"
	"github.com/mustanish/omelette/app/responses"
	"github.com/mustanish/omelette/app/schemas"
	"github.com/thedevsaddam/govalidator"
)

// ValidateBody validates request body
func ValidateBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var (
			route  = strings.TrimSuffix(req.URL.String(), "/")
			method = req.Method
		)
		if value, exist := schemas.Schema[route+":"+method]; exist {
			var (
				schema   = value.(map[string]interface{})
				rules    = govalidator.MapData{}
				messages = govalidator.MapData{}
			)

			for key, value := range schema["rules"].(map[string][]string) {
				rules[key] = value
			}

			for key, value := range schema["messages"].(map[string][]string) {
				messages[key] = value
			}
			var (
				opts = govalidator.Options{
					Data:     schema["data"],
					Request:  req,
					Rules:    rules,
					Messages: messages,
				}
				err = govalidator.New(opts).ValidateJSON()
			)
			if len(err) > 0 {
				if _, exist := err["_error"]; exist {
					render.Render(res, req, responses.NewHTTPError(http.StatusBadRequest, constants.InvalidReq))
				} else {
					render.Render(res, req, responses.NewHTTPError(http.StatusBadRequest, err))
				}
				return
			}
			ctx := context.WithValue(req.Context(), "data", schema["data"])
			next.ServeHTTP(res, req.WithContext(ctx))
		}
	})
}

// ValidateHeader validates request headers, params, query
func ValidateHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		next.ServeHTTP(res, req)
	})
}
