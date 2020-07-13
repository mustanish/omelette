package middlewares

import (
	"context"
	"net/http"
	"omelette/app/constants"
	"omelette/app/responses"
	"omelette/app/schemas"
	"strings"

	"github.com/go-chi/render"
	"github.com/thedevsaddam/govalidator"
)

// Validate validates the current request
func Validate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var (
			route  = strings.TrimSuffix(req.URL.String(), "/")
			method = strings.ToLower(req.Method)
		)
		opts, data := schemas.MapOpts(req, route, method)
		if opts != nil && data != nil {
			var (
				opts     = opts.(map[string]interface{})
				rules    = govalidator.MapData{}
				messages = govalidator.MapData{}
			)
			if opts["rules"] != nil {
				for key, value := range opts["rules"].(map[string][]string) {
					rules[key] = value
				}
			}
			if opts["messages"] != nil {
				for key, value := range opts["messages"].(map[string][]string) {
					messages[key] = value
				}
			}
			var (
				options = govalidator.Options{
					Data:     data,
					Request:  req,
					Rules:    rules,
					Messages: messages,
				}
				err = govalidator.New(options).ValidateJSON()
			)
			if len(err) > 0 {
				if _, exist := err["_error"]; exist {
					render.Render(res, req, responses.NewHTTPError(http.StatusBadRequest, constants.InvalidReq))
				} else {
					render.Render(res, req, responses.NewHTTPError(http.StatusBadRequest, err))
				}
				return
			}
			ctx := context.WithValue(req.Context(), "data", data)
			next.ServeHTTP(res, req.WithContext(ctx))
		}
	})
}
