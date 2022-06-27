package admincontrollers

import (
	//"fmt"

	"github.com/gofiber/fiber/v2"
	"main.go/database"
	"main.go/models"
	"main.go/models/adminData"
	
	
)

func GetAdmin(c *fiber.Ctx) error{
	db := database.DB
	var admin []models.User

	db.Find(&admin, "role = ?", "admin")
	return c.JSON((admin))
}

func UpdateAdmin(c *fiber.Ctx) error {
	var data adminData.AdminDetails

	id := c.Params("id")

	err := c.BodyParser(&data)
	// fmt.Println(data)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Review your inputs",
			"data":    err,
		})
	}
	
	var user adminData.AdminDetails
	database.DB.Find(&user, "id= ?", id)

	data = adminData.AdminDetails{
		Id:              user.Id,
		Name:            user.Name,
		Email:           user.Email,
		Telephone:       data.Telephone,
		Gender:          data.Gender,
		ProfileImageurl: data.ProfileImageurl,
	}

	database.DB.Save(&data)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Admin found",
		"error":   err,
		"data":    data,
	})
}

