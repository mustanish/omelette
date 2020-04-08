package tests_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/mustanish/omelette/app/routes"
	userschemas "github.com/mustanish/omelette/app/schemas/user"
	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"
)

var _ = Describe("Auth APIs", func() {
	Describe("POST /auth", func() {
		Context("when no identity is passed", func() {
			var data userschemas.Authenticate
			It("should fail", func() {
				data, _ := json.Marshal(data)
				req, _ := http.NewRequest("POST", "/auth", bytes.NewBuffer(data))
				req.Header.Set("Content-Type", "application/json")
				res := httptest.NewRecorder()
				routes.RouterInstance().ServeHTTP(res, req)
				log.Println(res)
			})
		})

		/*Context("when identity passed is of the wrong format", func() {
			It("should fail, when passed email is of the wrong format", func() {
				http.NewRequest("POST", "/auth")
			})
			It("should fail, when passed phone number is of the wrong format", func() {
				http.NewRequest("POST", "/auth")
			})
		})*/
	})
})
