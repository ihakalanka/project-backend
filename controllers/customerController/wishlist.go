package customerController

import (
	"github.com/gofiber/fiber/v2"
	
	"main.go/database"
	"main.go/models/customerData"
)

func Postlist(c *fiber.Ctx) error {
		db := database.DB
		var list customerData.WishlistData
		err := c.BodyParser(&list)
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  "error",
				"message": "Error",
				"data":    err,
			})
		}
		db.Create(&list)
		return c.JSON(list)
	}


func Deletelist(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	var list customerData.WishlistData
	err := db.Find(&list, "id = ?", id).Error
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  "error",
				"message": "error in delete list",
			})
		}
	db.Unscoped().Delete(&list)
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "list data deleted",
		})	
}
func GetlistbyUserId(c *fiber.Ctx) error {
	db := database.DB
	var list []customerData.WishlistData
	id := c.Params("id")
		err := db.Find(&list, "user_id = ?", id).Error
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  "error",
				"message": "error in get data in cart",
			})
		}
	return c.JSON(list)
}