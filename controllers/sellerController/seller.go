package sellercontroller

import (
	"reflect"
	"fmt"
	"strconv"
	"github.com/gofiber/fiber/v2"
	"main.go/database"
	"main.go/models"
	"main.go/models/sellerData"
	"main.go/models/adminData"
)
func GetAllSellers(c *fiber.Ctx) error{
	db:=database.DB
	var user []models.User
	
	
	db.Find(&user, "role = ?", "Seller")

	return c.JSON(user)
}
func GetSellerId(c *fiber.Ctx) error{
	db := database.DB
	id := c.Params("id")
	var seller models.User
	db.Find(&seller,id)
	return c.JSON(seller)
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
	//  err1 := strconv.ParseUint(id,10,32)
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
	user.Role = user1.Role
	// user.Previlage = user1.Previlage

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
func Count(c *fiber.Ctx) error{
	 db:=database.DB
	var user []models.User
	db.Find(&user)
	arrayLength:=len(user)
	var count models.Count
	count.CountUser= arrayLength
	// fmt.Println(count)
	return	c.JSON(count)
}
func CountProduct(c *fiber.Ctx) error{
	db:=database.DB
	var product []sellerData.Productdata
	db.Find(&product)
	arrayLength:=len(product)
	var count sellerData.ProductCount
	count.CountProduct=arrayLength
	// fmt.Println(count) 
	return c.JSON(count)
}
func CountSeller(c *fiber.Ctx) error{
	db:=database.DB
	var seller []models.User
	db.Find(&seller,"role = ?", "Seller")

	var customer []models.User
	db.Find(&customer,"role = ?", "Customer")

	arrayLength1:=len(seller)
	var sellercount models.Count
	sellercount.CountUser=arrayLength1

	arrayLength2:=len(customer)
	var customercount models.Count
	customercount.CountUser=arrayLength2

	// fmt.Println(sellercount,customercount)
	return c.JSON(fiber.Map{
		"sellercount":sellercount,
		"customercount":customercount,
	})
}

func GetCatWiseCount(c *fiber.Ctx) error {
	var category []adminData.Category
	database.DB.Find(&category)
 
	count := len(category)
	var prod []sellerData.Productdata
 
	type T struct {
		CategoryName string `json:"categoryName"`
		ProdCount    int    `json:"prodCount"`
	}
 
	var arr1 [5]T
	for i := 0; i < count; i++ {
		database.DB.Find(&prod, "category_name = ?", category[i].CategoryName)
 
		
			arr1[i].CategoryName = category[i].CategoryName
			arr1[i].ProdCount = len(prod)
		
 
	}
 
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":200,
		"data": arr1,
		
	})
}

func CountCategory(c *fiber.Ctx) error{
	db := database.DB
	var category []adminData.Category
	db.Find(&category)
	arrayLength := len(category)

	var count adminData.CategoryCount
	count.CountCategory=arrayLength
	fmt.Println(count)
	return c.JSON(count)
}

// func GetCatWiseProductCount(c *fiber.Ctx) error {
// 	var category []adminData.Category
// 	database.DB.Find(&category)
 
// 	count := len(category)
// 	var prod []sellerData.Productdata
 
// 	type T struct {
// 		CategoryName string `json:"categoryName"`
// 		ProdCount    int    `json:"prodCount"`
// 	}
	
// 	var arr1 [5]T
// 	for i := 0; i < count; i++ {
// 		database.DB.Find(&prod, "category_name = ?", category[i].CategoryName)
 
		
// 			arr1[i].CategoryName = category[i].CategoryName
// 			arr1[i].ProdCount += prod[i].Productquantity
		
 
// 	}
 
// 	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
// 		"status":200,
// 		"data": arr1,
		
// 	})
// }
