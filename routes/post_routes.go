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
			"password": newUser.Password,
		},
	})
}
