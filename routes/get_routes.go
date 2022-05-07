package routes

import (
	"fmt"
	"net/http"

	"github.com/ValonRexhepi/JWT-Authentication-System/controllers"
	"github.com/ValonRexhepi/JWT-Authentication-System/models"
	"github.com/gin-gonic/gin"
)

// AuthenticateUser function to respond to a jwt authentication of a user.
// Respond by JSON object with error if error, else respond with
// success message.
func AuthenticateUser(c *gin.Context) {

	cookie, err := c.Request.Cookie("token")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprint(err),
		})
		return
	}

	tokenString := cookie.Value
	claim := &models.Claim{}

	token, err := controllers.GetToken(tokenString, claim)

	if err != nil || !token.Valid {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprint(err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":  "Successfully connected",
		"username": claim.Username,
	})
}
