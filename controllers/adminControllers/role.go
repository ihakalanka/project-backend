package admincontrollers

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"fmt"
	"main.go/database"
	"main.go/models/adminData"
)

func Getrole(c *fiber.Ctx) error {
	db := database.DB
	var role []adminData.Role
	db.Find(&role)
	
	return c.JSON(role)
}

func Postrole(c *fiber.Ctx) error {
	db := database.DB
	var role2 adminData.Role
	var role adminData.Role
	// role := new(adminData.Role)
	 err := c.BodyParser(&role)
	 if err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Error",
			"data":    err,
		})
	}
	
	 err = db.Find(&role2, "role_name = ?", role.RoleName).Error
	 
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Error from method",
			"data":    err,
		})
	}

	if role2.RoleName == role.RoleName{
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Entered role is exist.",
		}) 
	}

	db.Create(&role)
	return c.JSON(role)
}

func Getroleid(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	var role adminData.Role
	db.Find(&role, id)
	return c.JSON(role)
}	

func Deleterole(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	var role adminData.Role
	err := db.Find(&role, "id = ?", id).Error
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  "error",
				"message": "error in delete role",
			})
		}
	db.Unscoped().Delete(&role)
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Role data deleted",
		})	
}

func Updaterole(c *fiber.Ctx) error {
	db := database.DB
	var role adminData.Role
	var role2 adminData.Role
	
	id := c.Params("id")

		err := c.BodyParser(&role)
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  "error",
				"message": "Review your inputs",
				"data":    err,
			})
		}	

		role.Id,err = strconv.Atoi(id)
		fmt.Println(err)
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  "error",
				"message": "Review your inputs",
				"data":    err,
			})
		}
		err = db.Find(&role2, "role_name = ?", role.RoleName).Error//*
	 
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Error from method",
			"data":    err,
		})
	}

		if role2.RoleName == role.RoleName{
			return c.JSON(fiber.Map{
				"status":  "error",
				"message": "Entered role is exist.",
			}) 
		}//*

		db.Save(&role)
			return c.JSON(fiber.Map{
			"status":  "success",
			"message": "role found",
			"error": err,
			"data":    role,
		})
}