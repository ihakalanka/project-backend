package merchantcontroller

import (
	"fmt"
	"strconv"
	"github.com/gofiber/fiber/v2"
	"main.go/database"
	"main.go/models/merchantData"
)

func PostMerchant(c *fiber.Ctx) error {
	db := database.DB
	merchant := new(merchantData.Merchantdata)
	if err := c.BodyParser(merchant); err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Error",
			"data":    err,
	})
}

	db.Create(&merchant)
	return c.JSON(merchant)
}

func GetMerchant(c *fiber.Ctx) error {
	db := database.DB
	var merchant []merchantData.Merchantdata
	db.Find(&merchant)

	return c.JSON(merchant)
}

func GetMerchantid(c *fiber.Ctx) error{
	db := database.DB
	id := c.Params("id")

	var merchant []merchantData.Merchantdata
	err := db.Find(&merchant,"userid = ?",id).Error
	if err != nil{
		return c.JSON(fiber.Map{
			"error":err,
		})
	}
	return c.JSON(merchant)
}

func UpdateMerchant(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	var merchant merchantData.Merchantdata
	err := c.BodyParser(&merchant)
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  "error",
				"message": "Review your inputs",
				"data":    err,
			})
		}
		merchant.Id,err = strconv.Atoi(id)
		fmt.Println(err)

		db.Save(&merchant)
			return c.JSON(fiber.Map{
			"status":  "success",
			"message": "updated profile",
			"error": err,
			"data": merchant,
		})
}
