# JWT Authentication System
## Goal
The goal of this project is to create a Go application that implements a JWT 
token Authentication. This API will serve as a demonstration and training for 
the use of various technologies with Go. The token is valid 15 minutes but 
can be refresh if 1 minute remains from is expiration.

## Methods Available
- Register an User
- Login an User
- Authenticate User with token

## Routes Requests
- [POST]
  - localhost:8080/users/register &rarr; Register an User
  - localhost:8080/users/login &rarr; Login an User and Create JWT Token
- [GET]
  - localhost:8080/users/login &rarr; Authenticate an User with JWT Token

## Libraries Used
- [Gorm, ORM Management Library](https://gorm.io/index.html)
- [Gin, HTTP WebFramework](https://github.com/gin-gonic/gin)
- [bcrypt, Password Encoder Decoder](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
- [go password validator, Password Entropy Validator](https://github.com/wagslane/go-password-validator)
- [golang-jwt v4](https://pkg.go.dev/github.com/golang-jwt/jwt/v4)

## Docker Image
- [MySql Docker Image :latest](https://hub.docker.com/_/mysql)
- User "root" / Password "secret" (if needed)

## How to Use ?
1. Launch the container with the command "docker-container up". 
2. Run the program in Go with the command "go run main.go". 
3. Query the webserver (for example with postman) for the different available requests.