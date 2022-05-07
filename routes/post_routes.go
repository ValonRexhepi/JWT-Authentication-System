package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ValonRexhepi/JWT-Authentication-System/controllers"
	"github.com/ValonRexhepi/JWT-Authentication-System/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// AddUser function to respond to a post a new user.
// Respond by JSON object with error if error,
// else respond with success message and username and mail of created user.
func AddUser(c *gin.Context) {
	var newUser models.User

	if err := c.Bind(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Error of type : %v", err),
		})
		return
	}

	id, err := controllers.AddUser(&newUser)

	if id == -1 || err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Error of type : %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "Successfully created user",
		"usercreated": map[string]string{
			"username": newUser.Username,
			"email":    newUser.Email,
		},
	})
}

// LoginUser function to respond to a login of a user.
// Respond by JSON object with error if error,
// else respond with success message and JWT cookie.
func LoginUser(c *gin.Context) {
	var loginUser models.User

	if err := c.Bind(&loginUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Error of type : %v", err),
		})
		return
	}

	err := controllers.LoginUser(&loginUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Error of type : %v", err),
		})
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	})

	tokenString, err := claims.SignedString([]byte(controllers.JwtSecretKey))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("something went wrong with the connection"),
		})
		return
	}

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}

	http.SetCookie(c.Writer, cookie)

	c.JSON(http.StatusOK, gin.H{
		"success": "Successfully connected",
	})
}
