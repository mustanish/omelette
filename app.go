package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	config "github.com/mustanish/omelette/configs"
	"github.com/mustanish/omelette/controllers"
	"github.com/mustanish/omelette/entity"
	"github.com/mustanish/omelette/helpers"
	response "github.com/mustanish/omelette/responses"
)

var (
	option        config.Option
	value         config.Value
	errorResponse response.Error
)

func init() {
	value = option.Load()
	entity.Mysql(value.Dbuser, value.Dbpassword, value.Dbhost, value.Dbname, value.Dbport)
}

func notFound(res http.ResponseWriter, req *http.Request) {
	errorResponse.Error = config.NotFound
	errorResponse.Code = http.StatusNotFound
	helpers.SetResponse(res, http.StatusNotFound, errorResponse)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user", helpers.Logger(controllers.Authorize)).Methods("POST")
	router.HandleFunc("/otp/verify", helpers.Logger(helpers.VerifyToken(controllers.Verify))).Methods("PATCH")
	router.HandleFunc("/user/{username}", helpers.Logger(helpers.VerifyToken(controllers.ReadDetail))).Methods("GET")
	router.HandleFunc("/user/{username}", helpers.Logger(helpers.VerifyToken(controllers.UpdateDetail))).Methods("PATCH")
	router.NotFoundHandler = http.HandlerFunc(notFound)
	log.Println("Servers started on: http://localhost:" + value.Port)
	http.ListenAndServe(":"+value.Port, helpers.RemoveTrailingSlash(router))
}
