package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"main.go/models"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:iha075@tcp(127.0.0.1:3306)/projectTest?charset=utf8mb4&parseTime=True&loc=Local"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
