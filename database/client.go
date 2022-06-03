package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"main.go/models"
	"main.go/models/adminData"
	"main.go/models/sellerData"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:Sehajini97.@tcp(127.0.0.1:3306)/test1?charset=utf8mb4&parseTime=True&loc=Local"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{}, &models.PasswordReset{})
	connection.AutoMigrate(&adminData.Category{})
	connection.AutoMigrate(&sellerData.Productdata{})
	connection.AutoMigrate(&adminData.Role{})
	
}

