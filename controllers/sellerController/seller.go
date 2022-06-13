package sellercontroller

import (
	"reflect"
	"fmt"
	"strconv"
	"github.com/gofiber/fiber/v2"
	"main.go/database"
	"main.go/models"
)
func GetAllSellers(c *fiber.Ctx) error{
	db:=database.DB
	var user []models.User
	
	
	db.Find(&user, "role = ?", "Seller")

	return c.JSON(user)
}
func DeleteSeller(c *fiber.Ctx) error{
	db := database.DB
	id := c.Params("id")
	var user models.User

	err := db.Find(&user, "id = ?", id).Error
	if err != nil{
		return c.JSON(fiber.Map{
			"status": "error",
			"message": "error in delete user",
		})
	}
	db.Unscoped().Delete(&user)
		return c.JSON(fiber.Map{
			"status": "success",
			"message": "User data deleted",
		})
}
func UpdateSeller(c *fiber.Ctx) error{
	db := database.DB
	var user models.User
	

	id := c.Params("id")

	err := c.BodyParser(&user)
	if err != nil{
		return c.JSON(fiber.Map{
			"status":"error",
			"message":"Check your inputs",
			"data":err,
		})
	}
var user1 models.User
err = db.Find(&user1,"id = ?",id).Error
if err != nil {
	return c.JSON(fiber.Map{
		"error":err,
	})
}

	u64,err1 := strconv.ParseUint(id,10,32)
		// fmt.Println(err)
		if err1 != nil {
			return c.JSON(fiber.Map{
				"status":  "error",
				"message": "Check your inputs",
				"data":    err1,
			})
		}
	user.Id=uint(u64)
	user.FirstName = user1.FirstName
	user.LastName = user1.LastName
	user.Email = user1.Email
	user.Password = user1.Password
	
	fmt.Println(reflect.TypeOf(user.Id))
	fmt.Println(user)
		db.Save(&user)
		return c.JSON(fiber.Map{
		"status":  "success",
		"message": "User found",
		"error": err,
		"data":    user,
	})
}
