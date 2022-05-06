package server

import (
	"fmt"
	"os"

	"github.com/ValonRexhepi/JWT-Authentication-System/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// LaunchServer function to define default router, define the routes
// and run the server on port localhost:8080.
func LaunchServer() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowAllOrigins = true

	router.Use(cors.New(config))

	router.POST("/users/register", routes.AddUser)
	router.POST("/users/login", routes.LoginUser)

	err := router.Run("localhost:8080")

	if err != nil {
		fmt.Println("Can't launch the server: ", err)
		os.Exit(-1)
	}
}
