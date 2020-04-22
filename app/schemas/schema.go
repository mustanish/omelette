package schemas

import (
	"net/http"

	"github.com/go-chi/chi"
	authschemas "github.com/mustanish/omelette/app/schemas/auth"
	bookschemas "github.com/mustanish/omelette/app/schemas/book"
	userschemas "github.com/mustanish/omelette/app/schemas/user"
)

//MapOpts is maps requested URL with corresponding validation options
func MapOpts(req *http.Request, route string, method string) (interface{}, interface{}) {
	schema := map[string]interface{}{
		"/auth:post":  authschemas.AuthenticateOpts,
		"/login:post": authschemas.LoginOpts,
		"/user:patch": userschemas.UpdateUserOpts,
		"/book:post":  bookschemas.AddBookOpts,
		"/book/" + chi.URLParam(req, "id") + ":patch": bookschemas.UpdateBookOpts,
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
		return new(authschemas.Authenticate)
	case "/login:post":
		return new(authschemas.Login)
	case "/user:patch":
		return new(userschemas.UpdateUser)
	case "/book:post":
		return new(bookschemas.AddBook)
	case "/book/" + chi.URLParam(req, "id") + ":patch":
		return new(bookschemas.UpdateBook)
	}
	return nil
}
