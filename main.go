package main

import (
	"github.com/ValonRexhepi/JWT-Authentication-System/controllers"
	"github.com/ValonRexhepi/JWT-Authentication-System/server"
)

// Main function of the application, connect to the database, migrate the
// schemas and launch the Web Server.
func main() {
	controllers.Connect()
	controllers.Migrate()
	server.LaunchServer()
}
