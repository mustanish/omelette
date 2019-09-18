package helpers

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type logger struct {
	RequestMethod   string      `json:"requestMethod"`
	RequestURL      string      `json:"requestUrl"`
	RequestBody     interface{} `json:"requestBody"`
	RequestHeaders  interface{} `json:"requestHeaders"`
	ResponseCode    int         `json:"responseCode"`
	ResponseHeaders interface{} `json:"responseHeaders"`
	ResponseBody    interface{} `json:"responseBody"`
}

// Logger is used to log every request being served
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		filePath, _ := filepath.Abs("logs/" + time.Now().Format("2006-01-02") + ".log")
		file, _ := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		defer file.Close()
		reqBody := make(map[string]interface{})
		resBody := make(map[string]interface{})
		body, _ := ioutil.ReadAll(req.Body)
		req.Body = ioutil.NopCloser(bytes.NewReader(body))
		rec := httptest.NewRecorder()
		next.ServeHTTP(rec, req)
		_ = json.Unmarshal(rec.Body.Bytes(), &resBody)
		_ = json.Unmarshal(body, &reqBody)
		logger := logger{req.Method, req.Host, reqBody, req.Header, rec.Code, rec.HeaderMap, resBody}
		for key, value := range rec.HeaderMap {
			res.Header()[key] = value
		}
		res.WriteHeader(rec.Code)
		rec.Body.WriteTo(res)
		jsonString, _ := json.Marshal(logger)
		log.SetOutput(file)
		log.Println(string(jsonString))
	})
}

// VerifyToken is used to validate token
func VerifyToken(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		next.ServeHTTP(res, req)
	})
}

// Recover is used to recover from panic situation
func Recover() {

}

// SendSms is used to send sms
func SendSms() {

}

// SendEmail is used to send email
func SendEmail() {

}

// FormatError is used format error
func FormatError(error string) map[string]string {
	errorMap := make(map[string]string)
	errors := strings.Split(error, ";")
	for _, value := range errors {
		indiErr := strings.Split(value, ":")
		errorMap[indiErr[0]] = indiErr[1]
	}
	return errorMap
}

// RemoveTrailingSlash is used to remove trailing slash from the url
func RemoveTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		req.URL.Path = strings.TrimSuffix(req.URL.Path, "/")
		next.ServeHTTP(res, req)
	})
}

// SetResponse is used to set response
func SetResponse(res http.ResponseWriter, status int, data interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	json.NewEncoder(res).Encode(data)
}

// GenerateOtp is used to generate otp
func GenerateOtp(max int) string {
	table := [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		log.Println(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

// MaskEmail is used to email
func MaskEmail(email string) string {
	split := strings.Split(email, "@")
	partOne := split[0]
	partTwo := split[1]
	partOne = email[0:(len(partOne) - len(partOne)/2)]
	return partOne + "****@" + partTwo
}

// MaskNumber is used to mask
func MaskNumber(number string) string {
	return strings.Repeat("*", len(number)-4) + number[len(number)-4:]
}
