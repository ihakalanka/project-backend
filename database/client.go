package database

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"main.go/models"
	"main.go/models/adminData"
	"main.go/models/sellerData"
	"main.go/models/customerData"
	"main.go/models/companyData"
	
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
	connection.AutoMigrate(&adminData.Category{})
	connection.AutoMigrate(&sellerData.Productdata{})
	connection.AutoMigrate(&adminData.Role{})
	connection.AutoMigrate(&customerData.Cart{})
	connection.AutoMigrate(&customerData.WishlistData{})
	connection.AutoMigrate(&CompanyData.CompanyData{})
	
}


