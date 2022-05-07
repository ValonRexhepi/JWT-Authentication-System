package routes

import (
	"fmt"
	"net/http"
	"time"

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
			"error": "cannot get token",
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
		"success":  "successfully connected",
		"username": claim.Username,
	})
}

// RefreshUser function to refresh the JWT token of the authenticated user.
// Only refresh the token if a maximum of 1 minute remains before expiration.
// Respond by JSON object with error if error, else set new cookie with
// refreshed expiration time.
func RefreshUser(c *gin.Context) {
	cookie, err := c.Request.Cookie("token")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot get token",
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

	if time.Until(time.Unix(claim.ExpiresAt.Unix(), 0)) > time.Minute {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot refresh token",
		})
		return
	}

	tokenString, expirationTime, err := controllers.RefreshTokenWithTime(claim,
		time.Now().Add(time.Minute*15))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot refresh token",
		})
		return
	}

	http.SetCookie(c.Writer, controllers.GetJWTCookie(tokenString,
		expirationTime))

	c.JSON(http.StatusOK, gin.H{
		"success": "successfully refreshed token",
	})
}

// Logout function to logout the user and destroy the JWT token.
// Respond by JSON object with error if error, else destroy token.
func Logout(c *gin.Context) {
	cookie, err := c.Request.Cookie("token")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot get token",
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

	// To destoy a token we need to define a expire time prior to now.
	tokenString, expirationTime, err := controllers.RefreshTokenWithTime(claim,
		time.Now().Add(-time.Hour*25))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot logout",
		})
		return
	}

	http.SetCookie(c.Writer, controllers.GetJWTCookie(tokenString,
		expirationTime))

	c.JSON(http.StatusOK, gin.H{
		"success": "successfully logout token",
	})
}
