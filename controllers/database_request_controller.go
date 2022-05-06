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

	if err := ValidatePasswordEntropy(userToAdd.Password); err != nil {
		return -1, err
	}

	cryptPassword, err := GenerateCryptPassword(userToAdd.Password)

	if err != nil {
		return -1, err
	}

	userToAdd.Password = string(cryptPassword)

	result := DB.Omit("ID").Create(userToAdd)

	if result.Error != nil {
		return -1, result.Error
	}

	return userToAdd.ID, nil
}
