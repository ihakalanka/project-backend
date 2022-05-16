package admincontrollers
import (
	"github.com/gofiber/fiber/v2"
	
	"main.go/database"
	"main.go/models/adminData"
)

func Get(c *fiber.Ctx) error {
	db := database.DB
	var category []adminData.Category
	db.Find(&category)
	
	return c.JSON(category)
}

func Post(c *fiber.Ctx) error {
		db := database.DB
		category := new(adminData.Category)
		if err := c.BodyParser(category); err != nil {
			return c.JSON(fiber.Map{
				"status":  "error",
				"message": "Error",
				"data":    err,
			})
		}
		db.Create(&category)
		return c.JSON(category)
	}
