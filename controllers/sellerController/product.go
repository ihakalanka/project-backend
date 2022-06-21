package sellercontroller
import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"fmt"
	"main.go/database"
	"main.go/models/sellerData"
)

func Postproduct(c *fiber.Ctx) error {
		db := database.DB
		product := new(sellerData.Productdata)
		err := c.BodyParser(product)
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  201,
				"message": "Error",
				"data":    err,
			})
		}
		db.Create(&product)
		return c.JSON(product)
	}

func Getproduct(c *fiber.Ctx) error {
	db := database.DB
	var product []sellerData.Productdata
	db.Find(&product)
	
	return c.JSON(product)
}

func Getproductid(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	var product sellerData.Productdata
	db.Find(&product, id)
	return c.JSON(product)
}	

func Deleteproduct(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	var product sellerData.Productdata
	err := db.Find(&product, "id = ?", id).Error
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  "error",
				"message": "error in delete Product",
			})
		}
	db.Unscoped().Delete(&product)
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Product data deleted",
		})	
}

func Updateproduct(c *fiber.Ctx) error {
	db := database.DB
	var product sellerData.Productdata
	id := c.Params("id")
	
		err := c.BodyParser(&product)
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  "error",
				"message": "Review your inputs",
				"data":    err,
			})
		}
		product.Id,err = strconv.Atoi(id)
		fmt.Println(err)

		db.Save(&product)
			return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Product found",
			"error": err,
			"data":  product,
		})
}
func GetProductByUserId(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	var product []sellerData.Productdata
	err := db.Find(&product, "user_id = ?", id).Error
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  "error",
				"message": "error in delete Product",
			})
		}
	return c.JSON(product)
}