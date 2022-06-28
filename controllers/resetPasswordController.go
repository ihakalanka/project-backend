package controllers

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"main.go/database"
	"main.go/models"
	"math/rand"
	"net/smtp"
)

func Forgot(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	if err := database.DB.Find(&user, "email = ?", data["email"]).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "Error",
			"message": "There is an error in finding email method",
		})
	}

	if user.Email == data["email"] {

		email := data["email"]
		token := RandStingRunes(100)

		passwordReset := models.PasswordReset{
			Email: email,
			Token: token,
		}

		database.DB.Create(&passwordReset)

		from := "admin@example.com"

		to := []string{
			email,
		}

		url := "http://localhost:3000/signin/resetpass/" + token

		message := []byte("Click <a href=\"" + url + "\">here</a> to reset your password")

		err := smtp.SendMail("0.0.0.0:1025", nil, from, to, message)

		if err != nil {
			return err
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "check your emails",
		})
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User does not exist",
		})
	}

}

func Reset(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		return c.Status(400).JSON(fiber.Map{
			"message": "password does not match",
		})
	}

	passwordReset := models.PasswordReset{}

	database.DB.Where("token = ?", data["token"]).Last(&passwordReset)

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	database.DB.Model(&models.User{}).Where("email = ?", passwordReset.Email).Update("password", password)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func VerifyMail(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	email := data["email"]

	passwordReset := models.PasswordReset{
		Email: email,
	}

	database.DB.Create(&passwordReset)

	from := "admin@example.com"

	to := []string{
		email,
	}

	url := "http://localhost:3000/signin/"

	message := []byte("Click <a href=\"" + url + "\">here</a> to verify your email")

	err := smtp.SendMail("0.0.0.0:1025", nil, from, to, message)

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "check your email",
	})
}

func RandStingRunes(n int) string {
	var letterRunes = []rune("0123456789abcdefghijklmnnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
