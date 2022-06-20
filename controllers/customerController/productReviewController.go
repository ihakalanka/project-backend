package customerController

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"main.go/database"
	"main.go/models"
	"main.go/models/customerData"
	"os"
	"strings"
)

func CreateReview(c *fiber.Ctx) error {
	var data customerData.Review

	if err := c.BodyParser(&data); err != nil {
		c.Status(fiber.StatusBadRequest)
		return err
	}

	token := c.Get("Authorization")

	tokenArray := strings.Split(token, "Bearer ")
	a := strings.Join(tokenArray, " ")
	to := strings.TrimSpace(a)

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(to, claims, keyFunc)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err,
		})
	}

	var user customerData.Review
	email, _ := claims["Email"]

	if err := database.DB.Find(&user, "user_email = ?", email).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "Error",
			"message": "There is an error in finding email method",
		})
	}

	if user.UserEmail == email && user.ProdId == data.ProdId {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "Error",
			"message": "You are already reviewed",
		})
	}

	var name models.User
	database.DB.Find(&name, "email = ?", email)

	review := customerData.Review{
		ProdId:    data.ProdId,
		Name:      name.FirstName + " " + name.LastName,
		Rating:    data.Rating,
		Comment:   data.Comment,
		UserEmail: email.(string),
	}

	database.DB.Create(&review)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":    review,
	})
}

func keyFunc(*jwt.Token) (interface{}, error) {
	SecretKey := os.Getenv("SECRETKEY")
	return []byte(SecretKey), nil
}

/*func ViewReview(c *fiber.Ctx) error {

}

func UpdateReview(c *fiber.Ctx) error {

}

func DeleteReview(c *fiber.Ctx) error {

}
*/
