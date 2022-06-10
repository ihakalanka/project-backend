package sellercontroller

import (
	"github.com/gofiber/fiber/v2"
	"main.go/database"
	"main.go/models"
)
func GetAllSellers(c *fiber.Ctx) error{
	db:=database.DB
	var users []models.User
	
	
	db.Find(&users, "role = ?", "Seller")

	return c.JSON(users)
}
