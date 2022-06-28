package addressControllers

import (

	"github.com/gofiber/fiber/v2"

	"main.go/database"
	"main.go/models/addressData"
)

func Getcat(c *fiber.Ctx) error {
	db := database.DB
	var address []AddressData.AddressData
	db.Find(&address)
	
	return c.JSON(address)
}

func Postcat(c *fiber.Ctx) error {
		db := database.DB
		var address  AddressData.AddressData

		err := c.BodyParser(&address); 
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  "error",
				"message": "Error",
				"data":    err,
			})
		}
		var address1 AddressData.AddressData
		name := address.ReceiverName
		err = db.Find(&address1, "receiver_name = ?", name).Error
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  "error",
				"message": "error in delete address",
			})
		}
		
		if name == address1.ReceiverName {
			return c.JSON(fiber.Map{
				"status":  "error",
				"message": "Duplicate role available for address",
			})
		}
		db.Create(&address)
		return c.JSON(address)
	}

