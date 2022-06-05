package database

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"main.go/models"
	"os"
)

var DB *gorm.DB

func Connect() {
	godotenv.Load()
	dsn := os.Getenv("DSN")
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{}, &models.PasswordReset{})
}
