package tests_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/mustanish/omelette/app/routes"
	userschemas "github.com/mustanish/omelette/app/schemas/user"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"syreclabs.com/go/faker"
)

var _ = Describe("Auth APIs", func() {
	Describe("POST /auth", func() {
		var data userschemas.Authenticate

		Context("when no identity is passed", func() {
			It("should fail, when neither email nor phone number is passed", func() {
				data, _ := json.Marshal(data)
				req, _ := http.NewRequest("POST", "/auth", bytes.NewBuffer(data))
				req.Header.Set("Content-Type", "application/json")
				res := httptest.NewRecorder()
				routes.RouterInstance().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusBadRequest))
			})
		})

		Context("when identity passed is of the wrong format", func() {
			It("should fail, when passed email is of the wrong format", func() {
				data.Identity = faker.RandomString(8)
				data, _ := json.Marshal(data)
				req, _ := http.NewRequest("POST", "/auth", bytes.NewBuffer(data))
				req.Header.Set("Content-Type", "application/json")
				res := httptest.NewRecorder()
				routes.RouterInstance().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusBadRequest))
			})

			It("should fail, when passed phone number is of the wrong format", func() {
				data.Identity = faker.Number().Number(8)
				data, _ := json.Marshal(data)
				req, _ := http.NewRequest("POST", "/auth", bytes.NewBuffer(data))
				req.Header.Set("Content-Type", "application/json")
				res := httptest.NewRecorder()
				routes.RouterInstance().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})
