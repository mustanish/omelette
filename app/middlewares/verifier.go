package middlewares

import (
	"context"
	"net/http"
	"omelette/app/constants"
	"omelette/app/responses"
	"omelette/helpers"
	"strings"

	"github.com/go-chi/render"
)

// VerifyToken verifies user token
func VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		bearerToken := strings.Split(req.Header.Get("Authorization"), " ")
		if len(bearerToken) != 2 {
			render.Render(res, req, responses.NewHTTPError(http.StatusUnauthorized, constants.InvalidToken))
			return
		}
		ID, tokenID, valid := helpers.VerifyToken(bearerToken[1])
		if !valid {
			render.Render(res, req, responses.NewHTTPError(http.StatusUnauthorized, constants.InvalidToken))
			return
		}
		ctx := context.WithValue(req.Context(), "ID", ID)
		ctx = context.WithValue(ctx, "tokenID", tokenID)
		next.ServeHTTP(res, req.WithContext(ctx))
	})
}
