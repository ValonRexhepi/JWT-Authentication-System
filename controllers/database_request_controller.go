package controllers

import (
	"fmt"

	"github.com/ValonRexhepi/JWT-Authentication-System/models"
)

// AddUser function to add a new user to the database.
// Returns the id of the new User and nil if no error else returns -1
// as id and the error.
func AddUser(userToAdd *models.User) (int, error) {
	userToAdd.TrimSpaces()
	containsEmptyField := userToAdd.ContainsEmptyField()

	if containsEmptyField {
		return -1, fmt.Errorf("you cannot have empty fields")
	}

	// We choose a entropy of 100, see password validator for information
	// on password entropy
	if err := ValidatePasswordEntropy(userToAdd.Password, 100); err != nil {
		return -1, err
	}

	cryptPassword, err := GenerateCryptPassword(userToAdd.Password)

	if err != nil {
		return -1, fmt.Errorf("cannot create user")
	}

	userToAdd.Password = string(cryptPassword)

	result := DB.Omit("ID").Create(userToAdd)

	if result.Error != nil {
		return -1, fmt.Errorf("cannot create user")
	}

	return userToAdd.ID, nil
}

// LoginUser function to respond to a login of a user.
// Returns username with nil error if the connection succeeded
// and empty string with error otherwise.
func LoginUser(userToLogin *models.User) (string, error) {
	userToLogin.TrimSpaces()
	var userInDatabase models.User

	result := DB.Where("Email = ?", userToLogin.Email).First(&userInDatabase)

	if result.Error != nil || userInDatabase.ID == 0 {
		return "", fmt.Errorf("wrong login information")
	}

	if err := CheckHashWithPassword(userInDatabase.Password,
		userToLogin.Password); err != nil {
		return "", fmt.Errorf("wrong login information")
	}

	return userInDatabase.Username, nil
}
