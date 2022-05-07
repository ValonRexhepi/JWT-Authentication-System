package controllers

import (
	passwordvalidator "github.com/wagslane/go-password-validator"
	"golang.org/x/crypto/bcrypt"
)

// String secret key used for jwt encryption. As the jwt documentation
// recommends "It is recommended to generate
// a key using crypto/rand or something equivalent. You need the same key
// for signing and validating.". In a production environment, we would
// generate a key using crypto/rand and would store it encrypted in a
// database but for the seek of the example, we use a simple string here.
const JwtSecretKey = "thisisasecretkey"

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
