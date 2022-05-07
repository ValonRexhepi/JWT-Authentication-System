package routes

import (
	"fmt"
	"net/http"

	"github.com/ValonRexhepi/JWT-Authentication-System/controllers"
	"github.com/ValonRexhepi/JWT-Authentication-System/models"
	"github.com/gin-gonic/gin"
)

// AddUser function to respond to a post a new user.
// Respond by JSON object with error if error,
// else respond with success message and username and mail of created user.
func AddUser(c *gin.Context) {
	var newUser models.User

	if err := c.Bind(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot add user",
		})
		return
	}

	id, err := controllers.AddUser(&newUser)
	if id == -1 || err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprint(err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "successfully created user",
		"usercreated": map[string]string{
			"username": newUser.Username,
			"email":    newUser.Email,
		},
	})
}

// LoginUser function to respond to a login of a user.
// Respond by JSON object with error if error, else respond with success
// message and JWT cookie.
func LoginUser(c *gin.Context) {
	var loginUser models.User

	if err := c.Bind(&loginUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot login user",
		})
		return
	}

	username, err := controllers.LoginUser(&loginUser)

	if err != nil || len(username) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprint(err),
		})
		return
	}

	tokenString, expirationTime, err := controllers.GenerateTokenString(username)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot login user",
		})
		return
	}

	http.SetCookie(c.Writer,
		controllers.GetJWTCookie(tokenString, expirationTime))

	c.JSON(http.StatusOK, gin.H{
		"success": "successfully connected",
	})
}
