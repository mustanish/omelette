package controllers

import (
	"log"
	"net/http"
)

// Authenticate is used validate email/mobile via OTP
func Authenticate(res http.ResponseWriter, req *http.Request) {
	log.Println("INSIDE Authenticate")
}

// VerifyUser is used to verify generated OTP
func VerifyUser(res http.ResponseWriter, req *http.Request) {
}

// UpdateUser is used to update details of existing user
func UpdateUser(res http.ResponseWriter, req *http.Request) {
}

// GetUser is used get details of existing user
func GetUser(res http.ResponseWriter, req *http.Request) {
}

// DeleteUser is used block existing user
func DeleteUser(res http.ResponseWriter, req *http.Request) {
}

func generateOTP() {

}

func generateToken() {

}
