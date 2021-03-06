package customerController

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"main.go/database"
	"main.go/models"
	"main.go/models/customerData"
	"main.go/models/sellerData"
	"os"
	"strconv"
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

	if err := database.DB.Find(&user, "user_email = ? AND prod_id = ?", email, data.ProdId).Error; err != nil {
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
		UserId:    int(name.Id),
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

func SetAverageRating(c *fiber.Ctx) error {
	var review []customerData.Review

	id := c.Params("id")

	database.DB.Find(&review, "prod_id = ?", id)

	length := len(review)

	sum := 0
	for i := 0; i < length; i++ {
		sum += review[i].Rating
	}

	var average int
	if sum == 0 {
		average = length
	} else {
		average = sum / length
	}

	var sell sellerData.Productdata

	database.DB.Find(&sell, "id = ?", id)

	averageRate := sellerData.Productdata{
		Id:              sell.Id,
		ProductTitle:    sell.ProductTitle,
		ProductSubtitle: sell.ProductSubtitle,
		CategoryName:    sell.CategoryName,
		Imageurl:        sell.Imageurl,
		Description:     sell.Description,
		Productprice:    sell.Productprice,
		Productquantity: sell.Productquantity,
		UserId:          sell.UserId,
		AverageRate:     average,
	}

	database.DB.Save(&averageRate)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"noOfUsers":     length,
		"sumOfRating":   sum,
		"averageRating": average,
		"message":       "success",
	})
}

func ViewAllReviewsByItem(c *fiber.Ctx) error {
	var review []customerData.Review

	id := c.Params("id")

	database.DB.Find(&review, "prod_id = ?", id)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":    review,
		"message": "success",
	})
}

func DeleteReview(c *fiber.Ctx) error {
	var review customerData.Review

	id := c.Params("id")

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

	if err := database.DB.Find(&user, "user_email = ? AND prod_id = ?", email, id).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "Error",
			"message": "There is an error in finding email method",
		})
	}

	fmt.Println(id, user.ProdId, user.UserEmail, email)

	err = database.DB.Find(&review, "prod_id = ?", id).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": "error in delete review",
		})
	}

	id1, _ := strconv.Atoi(id)

	if id1 == user.ProdId && review.UserEmail == user.UserEmail {
		database.DB.Unscoped().Delete(&review)
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"data":    err,
			"message": "error",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  201,
		"message": "review deleted successfully",
	})
}
