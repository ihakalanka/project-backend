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
	
	return c.JSON(fiber.Map{
				"status":  200,
				"message": "category add succesfully",
				"data":category,
			})
}

func Postcat(c *fiber.Ctx) error {
		db := database.DB
		var category  adminData.Category

		err := c.BodyParser(&category); 
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  "error",
				"message": "Error",
				"data":    err,
			})
		}
		var category1 adminData.Category
		name := category.CategoryName
		err = db.Find(&category1, "category_name = ?", name).Error
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  404,
				"message": "error in  category",
			})
		}
		
		if name == category1.CategoryName {
			return c.JSON(fiber.Map{
				"status":  404,
				"message": "Duplicate category available ",
			})
		}
		db.Create(&category)
		return c.JSON(fiber.Map{
				"status":  200,
				"message": "category add successfull",
				"data":category,
			})
	}

func Getcatid(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	var category adminData.Category
	db.Find(&category, id)
	return c.JSON(fiber.Map{
				"status":  200,
				"message": "category available",
				"data":category,
			})
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
		var category1 adminData.Category
		name := category.CategoryName
		err = db.Find(&category1, "category_name = ?", name).Error
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  404,
				"message": "error in  category",
			})
		}
		
		if name == category1.CategoryName {
			return c.JSON(fiber.Map{
				"status":  404,
				"message": "Can not update ,Duplicate category available ",
			})
		}
		category.Id,err = strconv.Atoi(id)
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  404,
				"message": "Review your inputs",
				"data":    err,
			})
		}
		fmt.Println(err)

		db.Save(&category)
			return c.JSON(fiber.Map{
			"status":  200,
			"message": "category updated",
			"error": err,
			"data":  category,
		})
}