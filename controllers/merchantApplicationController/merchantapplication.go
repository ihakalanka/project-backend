package merchantapplicationcontroller

import (
	"github.com/gofiber/fiber/v2"
	"main.go/database"
	"main.go/models/merchantApplicationData"
)

 
func GetMerchantApplication(c *fiber.Ctx) error {
	db := database.DB
	var merchant []merchantApplicationData.MerchantApplicationdata
	db.Find(&merchant)

	return c.JSON(merchant)
}

 
