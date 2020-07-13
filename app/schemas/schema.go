package schemas

import (
	"net/http"
	"omelette/app/schemas/validation"

	"github.com/go-chi/chi"
)

//MapOpts is maps requested URL with corresponding validation options
func MapOpts(req *http.Request, route string, method string) (interface{}, interface{}) {
	schema := map[string]interface{}{
		"/auth:post":  validation.AuthenticateOpts,
		"/login:post": validation.LoginOpts,
		"/user:patch": validation.UpdateUserOpts,
		"/book:post":  validation.AddBookOpts,
		"/book/" + chi.URLParam(req, "id") + ":patch": validation.UpdateBookOpts,
	}
	options, exist := schema[route+":"+method]
	if exist {
		data := mapReqStruct(req, route, method)
		return options, data
	}
	return nil, nil
}

func mapReqStruct(req *http.Request, route string, method string) interface{} {
	switch route + ":" + method {
	case "/auth:post":
		return new(validation.Authenticate)
	case "/login:post":
		return new(validation.Login)
	case "/user:patch":
		return new(validation.UpdateUser)
	case "/book:post":
		return new(validation.AddBook)
	case "/book/" + chi.URLParam(req, "id") + ":patch":
		return new(validation.UpdateBook)
	}
	return nil
}
