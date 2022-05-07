package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ValonRexhepi/JWT-Authentication-System/models"
	"github.com/golang-jwt/jwt/v4"
	passwordvalidator "github.com/wagslane/go-password-validator"
	"golang.org/x/crypto/bcrypt"
)

// Byte array secret key used for jwt encryption. As the jwt documentation
// recommends "It is recommended to generate
// a key using crypto/rand or something equivalent. You need the same key
// for signing and validating.". In a production environment, we would
// generate a key using crypto/rand and would store it encrypted in a
// database but for the seek of the example, we use a simple string here.
var JwtSecretKey = []byte("thisisasecretkey")

// GenerateCryptPassword function to take a string password
// and if error return empty string with error else
// return string hashed password and nil error.
func GenerateCryptPassword(password string) (string, error) {
	cryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), 15)

	if err != nil {
		return "", err
	}

	return string(cryptPassword), nil
}

// GenerateTokenString function to create a string token with 15 minutes
// expiration time and username of user.
// Return the string token, expiration time and nil if no error, else
// error not nil.
func GenerateTokenString(username string) (string, time.Time, error) {
	expirationTime := time.Now().Add(time.Minute * 15)

	claim := &models.Claim{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString(JwtSecretKey)

	return tokenString, expirationTime, err
}

func GetToken(tokenString string, claim *models.Claim) (*jwt.Token, error) {
	tkn, err := jwt.ParseWithClaims(tokenString, claim,
		func(token *jwt.Token) (interface{}, error) {
			if token.Method.Alg() != "HS256" {
				return nil, fmt.Errorf("cannot sign in with this token")
			}
			return JwtSecretKey, nil
		})
	return tkn, err
}

// GetJWTCookie function to create a http cookie for the given token string
// and expiration time. Return a reference to the cookie object.
func GetJWTCookie(tokenString string, expirationTime time.Time) *http.Cookie {
	return &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    tokenString,
		Expires:  expirationTime,
		HttpOnly: true,
	}
}

// CheckHashWithPassword function to check if a given password match
// the hash in the database. Return nil if match else return the err.
func CheckHashWithPassword(passwordHashed, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHashed),
		[]byte(password))
}

// ValidatePasswordEntropy checks the password entropy return
// error with nil if has enough entropy else return a description
// message on how to improve password entropy.
func ValidatePasswordEntropy(password string) error {
	return passwordvalidator.Validate(password, 100)
}
