package customerController

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	
	"main.go/database"
	"main.go/models/customerData"
)

func Postcart(c *fiber.Ctx) error {
		db := database.DB
		var cart customerData.Cart
		err := c.BodyParser(&cart)
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  "error",
				"message": "Error",
				"data":    err,
			})
		}
		db.Create(&cart)
		return c.JSON(cart)
	}

func Updatecart(c *fiber.Ctx) error {
	db := database.DB
	var cart customerData.Cart
	id := c.Params("id")
	
		err := c.BodyParser(&cart)
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  "error",
				"message": "Review your inputs",
				"data":    err,
			})
		}
		cart.Id,err = strconv.Atoi(id)
		

		db.Save(&cart)
			return c.JSON(fiber.Map{
			"status":  "success",
			"message": "cart found",
			"error": err,
			"data":  cart,
		})
}
func Deletecart(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	var cart customerData.Cart
	err := db.Find(&cart, "id = ?", id).Error
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  "error",
				"message": "error in delete cart",
			})
		}
	db.Unscoped().Delete(&cart)
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Product data deleted",
		})	
}
func Getcart(c *fiber.Ctx) error {
	db := database.DB
	var cart []customerData.Cart
	id := c.Params("id")
		err := db.Find(&cart, "user_id = ?", id).Error
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  "error",
				"message": "error in get data in cart",
			})
		}
	
	return c.JSON(cart)
}

