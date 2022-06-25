package controllers

import (
	"github.com/gofiber/fiber/v2"
	"main.go/database"
	"main.go/models"
	"main.go/models/sellerData"
)

func GetAllSearchResult(c *fiber.Ctx) error{
	var items []sellerData.Productdata

	var catName models.SearchProd

	err := c.BodyParser(&catName)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	database.DB.Find(&items, "category_name = ?", catName.UserSearchCat)
	return c.JSON(items)
}