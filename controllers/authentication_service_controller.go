package controllers

import (
	passwordvalidator "github.com/wagslane/go-password-validator"
	"golang.org/x/crypto/bcrypt"
)

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

// ValidatePasswordEntropy checks the password entropy return
// error with nil if has enough entropy else return a description
// message on how to improve password entropy.
func ValidatePasswordEntropy(password string) error {
	return passwordvalidator.Validate(password, 100)
}
