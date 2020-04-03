package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-chi/render"
	"github.com/mustanish/omelette/app/constants"
	"github.com/mustanish/omelette/app/helpers"
	"github.com/mustanish/omelette/app/responses"
)

// VerifyToken verifies user token
func VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		bearerToken := strings.Split(req.Header.Get("Authorization"), " ")
		if len(bearerToken) != 2 {
			render.Render(res, req, responses.NewHTTPError(http.StatusUnauthorized, constants.InvalidToken))
			return
		}
		ID, err := helpers.VerifyToken(bearerToken[1])
		if err != nil {
			render.Render(res, req, responses.NewHTTPError(http.StatusUnauthorized, constants.InvalidToken))
			return
		}
		ctx := context.WithValue(req.Context(), "ID", ID)
		next.ServeHTTP(res, req.WithContext(ctx))
	})
}
