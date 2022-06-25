package sellercontroller

import(
	"github.com/gofiber/fiber/v2"
	// "strconv"
	// "fmt"
	"main.go/database"
	"main.go/models/sellerData"
)

func GetproductCount(c *fiber.Ctx) error {
	db := database.DB
	id:=c.Params("id")
	var product []sellerData.Productdata
	type CountProduct struct{
		Count int `json:"count"`
	}
	err:= db.Find(&product, "user_id = ?",id)
	if err != nil {
			return c.JSON(fiber.Map{
				"status":  404,
				"message": "error in delete Product",
			})
		}
	var count CountProduct

	count.Count = len(product)
	
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "error in delete Product",
		"data": count,
	})
}