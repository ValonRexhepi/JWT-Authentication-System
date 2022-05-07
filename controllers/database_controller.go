package controllers

import (
	"log"

	"github.com/ValonRexhepi/JWT-Authentication-System/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB variable that contains reference to the gorm.db object.
var DB *gorm.DB

// Connect method that will initialize the connection
// to the database.
func Connect() {
	cst := "root:secret@tcp(127.0.0.1:3306)/UserDB?charset=utf8mb4&parseTime=True&loc=Local"
	connection, err := gorm.Open(mysql.Open(cst), &gorm.Config{})
	if err != nil {
		log.Fatalln("connection to database failed: ", err)
	}
	DB = connection
}

// Migrate method that will migrate the schemas
// to the database.
func Migrate() {
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalln("schema migration failed: ", err)
	}
}
