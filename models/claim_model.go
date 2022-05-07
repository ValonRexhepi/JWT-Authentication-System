package models

import (
	"github.com/golang-jwt/jwt/v4"
)

// User struct that represents a Claim.
type Claim struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
