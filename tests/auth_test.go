package tests_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"omelette/app/constants"
	"omelette/app/responses"
	"omelette/app/routes"
	"omelette/app/schemas/validation"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"syreclabs.com/go/faker"
)

var _ = Describe("Auth APIs", func() {
	var error responses.HTTPError
	var success responses.HTTPSucess
	Describe("POST /auth", func() {
		var data validation.Authenticate

		Context("when no identity is passed", func() {
			It("should fail, when neither email nor phone number is passed", func() {
				data, _ := json.Marshal(data)
				req, _ := http.NewRequest("POST", "/auth", bytes.NewBuffer(data))
				req.Header.Set("Content-Type", "application/json")
				res := httptest.NewRecorder()
				routes.RouterInstance().ServeHTTP(res, req)
				if err := json.NewDecoder(res.Body).Decode(&error); err != nil {
					log.Println("Unable to decode body because of ", err.Error())
				}
				Expect(res.Code).To(Equal(http.StatusBadRequest))
				Expect(error.Status).To(Equal("failed"))
				Expect(error.Error).ShouldNot(BeEmpty())
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
				if err := json.NewDecoder(res.Body).Decode(&error); err != nil {
					log.Println("Unable to decode body because of ", err.Error())
				}
				Expect(res.Code).To(Equal(http.StatusBadRequest))
				Expect(error.Status).To(Equal("failed"))
				Expect(error.Error).ShouldNot(BeEmpty())
			})

			It("should fail, when passed phone number is of the wrong format", func() {
				data.Identity = faker.Number().Number(8)
				data, _ := json.Marshal(data)
				req, _ := http.NewRequest("POST", "/auth", bytes.NewBuffer(data))
				req.Header.Set("Content-Type", "application/json")
				res := httptest.NewRecorder()
				routes.RouterInstance().ServeHTTP(res, req)
				if err := json.NewDecoder(res.Body).Decode(&error); err != nil {
					log.Println("Unable to decode body because of ", err.Error())
				}
				Expect(res.Code).To(Equal(http.StatusBadRequest))
				Expect(error.Status).To(Equal("failed"))
				Expect(error.Error).ShouldNot(BeEmpty())
			})
		})

		Context("when identity passed is of the correct format", func() {
			It("should pass, when when passed email is of the correct format", func() {
				data.Identity = faker.Internet().Email()
				data, _ := json.Marshal(data)
				req, _ := http.NewRequest("POST", "/auth", bytes.NewBuffer(data))
				req.Header.Set("Content-Type", "application/json")
				res := httptest.NewRecorder()
				routes.RouterInstance().ServeHTTP(res, req)
				if err := json.NewDecoder(res.Body).Decode(&success); err != nil {
					log.Println("Unable to decode body because of ", err.Error())
				}
				Expect(res.Code).To(Equal(http.StatusOK))
				Expect(success.Status).To(Equal("success"))
				Expect(success.Data).ShouldNot(BeEmpty())
			})

			It("should pass, when when passed phone number is of the correct format", func() {
				data.Identity = faker.Number().Number(10)
				data, _ := json.Marshal(data)
				req, _ := http.NewRequest("POST", "/auth", bytes.NewBuffer(data))
				req.Header.Set("Content-Type", "application/json")
				res := httptest.NewRecorder()
				routes.RouterInstance().ServeHTTP(res, req)
				if err := json.NewDecoder(res.Body).Decode(&success); err != nil {
					log.Println("Unable to decode body because of ", err.Error())
				}
				resData := success.Data.(map[string]interface{})
				accessToken = resData["accessToken"].(string)
				Expect(res.Code).To(Equal(http.StatusOK))
				Expect(success.Status).To(Equal("success"))
				Expect(success.Data).ShouldNot(BeEmpty())
			})
		})
	})

	Describe("POST /login", func() {
		var data validation.Login
		Context("when no access token or wrong access token is passed", func() {
			It("should fail, when no access token is passed", func() {
				data.OTP = constants.OTPTest
				data, _ := json.Marshal(data)
				req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(data))
				req.Header.Set("Content-Type", "application/json")
				res := httptest.NewRecorder()
				routes.RouterInstance().ServeHTTP(res, req)
				if err := json.NewDecoder(res.Body).Decode(&error); err != nil {
					log.Println("Unable to decode body because of ", err.Error())
				}
				Expect(res.Code).To(Equal(http.StatusUnauthorized))
				Expect(error.Status).To(Equal("failed"))
				Expect(error.Error).ShouldNot(BeEmpty())
			})

			It("should fail, when wrong access token is passed", func() {
				data.OTP = constants.OTPTest
				data, _ := json.Marshal(data)
				req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(data))
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Authorization", faker.RandomString(10))
				res := httptest.NewRecorder()
				routes.RouterInstance().ServeHTTP(res, req)
				if err := json.NewDecoder(res.Body).Decode(&error); err != nil {
					log.Println("Unable to decode body because of ", err.Error())
				}
				Expect(res.Code).To(Equal(http.StatusUnauthorized))
				Expect(error.Status).To(Equal("failed"))
				Expect(error.Error).ShouldNot(BeEmpty())
			})
		})

		Context("when no OTP or wrong OTP is passed", func() {
			It("should fail, when no OTP is passed", func() {
				data.OTP = ""
				data, _ := json.Marshal(data)
				req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(data))
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Authorization", "Bearer "+accessToken)
				res := httptest.NewRecorder()
				routes.RouterInstance().ServeHTTP(res, req)
				if err := json.NewDecoder(res.Body).Decode(&error); err != nil {
					log.Println("Unable to decode body because of ", err.Error())
				}
				Expect(res.Code).To(Equal(http.StatusBadRequest))
				Expect(error.Status).To(Equal("failed"))
				Expect(error.Error).ShouldNot(BeEmpty())
			})

			It("should fail, when wrong OTP is passed", func() {
				data.OTP = faker.Number().Number(6)
				data, _ := json.Marshal(data)
				req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(data))
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Authorization", "Bearer "+accessToken)
				res := httptest.NewRecorder()
				routes.RouterInstance().ServeHTTP(res, req)
				if err := json.NewDecoder(res.Body).Decode(&error); err != nil {
					log.Println("Unable to decode body because of ", err.Error())
				}
				Expect(res.Code).To(Equal(http.StatusBadRequest))
				Expect(error.Status).To(Equal("failed"))
				Expect(error.Error).ShouldNot(BeEmpty())
			})
		})

		Context("when correct OTP is passed", func() {
			It("should pass, when correct OTP is passed", func() {
				data.OTP = constants.OTPTest
				data, _ := json.Marshal(data)
				req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(data))
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Authorization", "Bearer "+accessToken)
				res := httptest.NewRecorder()
				routes.RouterInstance().ServeHTTP(res, req)
				if err := json.NewDecoder(res.Body).Decode(&success); err != nil {
					log.Println("Unable to decode body because of ", err.Error())
				}
				Expect(res.Code).To(Equal(http.StatusOK))
				Expect(success.Status).To(Equal("success"))
				Expect(success.Data).ShouldNot(BeEmpty())
			})
		})
	})
})
