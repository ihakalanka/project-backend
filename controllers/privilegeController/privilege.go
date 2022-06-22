package privilegecontroller

import (
	"fmt"
	"strconv"
	"github.com/gofiber/fiber/v2"
	"main.go/database"
	"main.go/models/privilegeData"
)

func PostPrivilege(c *fiber.Ctx) error {
	db := database.DB
	privilege := new(privilegeData.Privilegedata)
	if err := c.BodyParser(privilege); err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Error",
			"data":    err,
	})
}

	db.Create(&privilege)
	return c.JSON(privilege )
}

func GetPrivilege(c *fiber.Ctx) error {
	db := database.DB
	var privilege  []privilegeData.Privilegedata
	db.Find(&privilege)

	return c.JSON(privilege)
}

func GetPrivilegeid(c *fiber.Ctx) error{
	db := database.DB
	id := c.Params("id")

	var privilege []privilegeData.Privilegedata
	err := db.Find(&privilege,"userid = ?",id).Error
	if err != nil{
		return c.JSON(fiber.Map{
			"error":err,
		})
	}
	return c.JSON(privilege)
}

func DeletePrivilege(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	var privilege privilegeData.Privilegedata
	err := db.Find(&privilege, "id = ?", id).Error
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  "error",
				"message": "error in delete Privilege",
			})
		}
	db.Unscoped().Delete(&privilege)
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Privilege data deleted",
		})	
}


func UpdatePrivilege(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	var privilege privilegeData.Privilegedata
	err := c.BodyParser(&privilege)
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  "error",
				"message": "Review your inputs",
				"data":    err,
			})
		}
		privilege.Id,err = strconv.Atoi(id)
		fmt.Println(err)

		db.Save(&privilege)
			return c.JSON(fiber.Map{
			"status":  "success",
			"message": "updated profile",
			"error": err,
			"data": privilege,
		})
}
