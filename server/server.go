package server

import (
	"fmt"
	"os"

	"github.com/ValonRexhepi/JWT-Authentication-System/routes"
	"github.com/gin-gonic/gin"
)

// LaunchServer function to define default router, define the routes
// and run the server on port localhost:8080.
func LaunchServer() {
	router := gin.Default()

	router.POST("/users", routes.AddUser)

	err := router.Run("localhost:8080")
	if err != nil {
		fmt.Println("Can't launch the server: ", err)
		os.Exit(-1)
	}
}
