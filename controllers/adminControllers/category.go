package admincontrollers

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"fmt"
	"main.go/database"
	"main.go/models/adminData"
)

func Getcat(c *fiber.Ctx) error {
	db := database.DB
	var category []adminData.Category
	db.Find(&category)
	
	return c.JSON(category)
}

func Postcat(c *fiber.Ctx) error {
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

func Getcatid(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params(":id")
	var category adminData.Category
	db.Find(&category, id)
	return c.JSON(category)
}	

func Deletecat(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	var category adminData.Category
	err := db.Find(&category, "id = ?", id).Error
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  "error",
				"message": "error in delete category",
			})
		}
	db.Unscoped().Delete(&category)
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Category data deleted",
		})	
}
func Updatecat(c *fiber.Ctx) error {
	db := database.DB
	var category adminData.Category
	id := c.Params("id")
		
		err := c.BodyParser(&category)
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  "error",
				"message": "Review your inputs",
				"data":    err,
			})
		}
		category.Id,err = strconv.Atoi(id)
		fmt.Println(err)

		db.Save(&category)
			return c.JSON(fiber.Map{
			"status":  "success",
			"message": "category found",
			"error": err,
			"data":    category,
		})
}