# JWT Authentication System
## Goal

## Methods Available

## Routes Requests
- [POST]
- [GET]

## Libraries Used
- [Gorm, ORM Management Library](https://gorm.io/index.html)
- [Gin, HTTP WebFramework](https://github.com/gin-gonic/gin)
- [bcrypt, Password Encoder Decoder](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
- [go password validator, Password Entropy Validator](https://github.com/wagslane/go-password-validator)

## Docker Image
- [MySql Docker Image :latest](https://hub.docker.com/_/mysql)
- User "root" / Password "secret" (if needed)

## How to Use ?
1. Launch the container with the command "docker-container up". 
2. Run the program in Go with the command "go run main.go". 
3. Query the webserver (for example with postman) for the different available requests.