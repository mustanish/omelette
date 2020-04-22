package helpers

import (
	"crypto/rand"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mustanish/omelette/app/constants"
	tokens "github.com/mustanish/omelette/app/repository"
	uuid "github.com/satori/go.uuid"
)

// Claims represents the StandardClaims
type Claims struct {
	TokenID string `json:"tokenId"`
	jwt.StandardClaims
}

// GenerateToken generates access token
func GenerateToken(ID string, validity time.Duration) (string, string, int64, error) {
	var (
		now       = time.Now()
		expiresAt = now.Add(time.Second * validity)
		tokenID   = uuid.NewV4().String()
	)
	claims := &Claims{
		TokenID: tokenID,
		StandardClaims: jwt.StandardClaims{
			Id:        ID,
			IssuedAt:  now.Unix(),
			ExpiresAt: expiresAt.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(constants.Jwtsecret))
	return tokenString, tokenID, expiresAt.Unix(), err
}

// VerifyToken is used to validate token
func VerifyToken(token string) (string, string, bool) {
	var (
		userID  string
		tokenID string
		claims  = &Claims{}
	)
	decoded, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.Jwtsecret), nil
	})
	exist, _ := tokens.Exist(claims.TokenID)
	if err != nil || !decoded.Valid || !exist {
		return userID, tokenID, false
	}
	return claims.Id, claims.TokenID, true
}

// MaskEmail masks the provided email
func MaskEmail(email string) string {
	split := strings.Split(email, "@")
	return email[0:(len(split[0])-len(split[0])/2)] + "****@" + split[1]
}

// MaskNumber masks the provided phone number
func MaskNumber(number string) string {
	return strings.Repeat("*", len(number)-4) + number[len(number)-4:]
}

// GenerateOTP generates an OTP to verify email/phone
func GenerateOTP(length int) string {
	activeProfile := os.Getenv("ENV")
	if activeProfile == "testing" || activeProfile == "development" {
		return constants.OTPTest
	}
	table := [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	buffer := make([]byte, length)
	io.ReadAtLeast(rand.Reader, buffer, length)
	for i := 0; i < len(buffer); i++ {
		buffer[i] = table[int(buffer[i])%len(table)]
	}
	return string(buffer)
}

// SendEmail sends email to the provided email
func SendEmail() {

}

// sendSms sends sms to the provided phone number
func sendSms() {

}

// InArray Checks if a value exists within a slice
func InArray(needle interface{}, haystack interface{}) bool {
	if reflect.TypeOf(haystack).Kind() == reflect.Slice {
		j := reflect.ValueOf(haystack)
		for i := 0; i < j.Len(); i++ {
			if reflect.DeepEqual(needle, j.Index(i).Interface()) == true {
				return true
			}
		}
	}
	return false
}

// QueryParam returns query string from current request
func QueryParam(req *http.Request, key string) string {
	query := req.URL.Query()
	return query.Get(key)
}
