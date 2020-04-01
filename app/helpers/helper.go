package helpers

import (
	"crypto/rand"
	"io"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mustanish/omelette/app/constants"
)

// Claims represents the StandardClaims
type Claims struct {
	jwt.StandardClaims
}

// GenerateToken generates access token
func GenerateToken(ID string, validity time.Duration) (string, error) {
	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			Id:        ID,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Second * validity).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(constants.Jwtsecret))
	return tokenString, err
}

// VerifyToken is used to validate token
func VerifyToken(token string) (string, error) {
	var (
		ID     string
		claims = &Claims{}
	)
	decoded, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.Jwtsecret), nil
	})
	if err != nil || !decoded.Valid {
		return ID, err
	}
	return claims.Id, nil
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
