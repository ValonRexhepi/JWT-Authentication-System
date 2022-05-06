package models

import (
	"fmt"
	"strings"
)

// User struct that represents an User.
type User struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name     string `json:"name" gorm:"not null"`
	Username string `json:"username" gorm:"not null;unique"`
	Email    string `json:"email" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
}

// String method to return a string representation of the user.
func (u *User) String() string {
	return fmt.Sprintf("{%v, %v, %v, %v, %v}", u.ID, u.Name, u.Email,
		u.Username, u.Password)
}

// ContainsEmptyField method to check if name, username, email or password fields
// are empty.
// Return true if at least one is empty, otherwise return false.
func (u *User) ContainsEmptyField() bool {
	if len(strings.TrimSpace(u.Name)) == 0 || len(u.Email) == 0 ||
		len(u.Username) == 0 || len(strings.TrimSpace(u.Password)) == 0 {
		return true
	}

	return false
}

// TrimSpaces method to remove spaces from username and email.
func (u *User) TrimSpaces() {
	u.Username = strings.TrimSpace(u.Username)
	u.Email = strings.TrimSpace(u.Email)
}
