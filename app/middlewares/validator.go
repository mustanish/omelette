package middlewares

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/mustanish/omelette/app/constants"
	"github.com/mustanish/omelette/app/responses"
	"github.com/mustanish/omelette/app/schemas"
	"github.com/thedevsaddam/govalidator"
)

// Body is exposed to validate request body, headers, params, query
func Body(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var (
			route  = req.URL.String()
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
				if value, exist := err["_error"]; exist && value[0] == "EOF" {
					render.Render(res, req, responses.NewHTTPError(http.StatusBadRequest, constants.InvalidReq))
				} else {
					render.Render(res, req, responses.NewHTTPError(http.StatusBadRequest, err))
				}
				return
			}
			next.ServeHTTP(res, req)
		}
	})
}
