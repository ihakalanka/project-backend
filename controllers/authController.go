package controllers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"main.go/database"
	"main.go/models"
	"os"
	"strings"
	"time"
	"unicode"
)

func Register(c *fiber.Ctx) error {
	var data models.User
	if err := c.BodyParser(&data); err != nil {
		c.Status(fiber.StatusBadRequest)
		return err
	}

	if err := verifyPassword(data.Password); err != nil {
		c.Status(fiber.StatusBadRequest)
		return err
	}

	var user1 models.User
	email := data.Email

	if err := database.DB.Find(&user1, "email = ?", email).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "Error",
			"message": "There is an error in finding email method",
		})
	}

	if user1.Email == email {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "Error",
			"message": "You are already used this email address",
		})
	}

	cost := 14
	password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), cost)
	user := models.User{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Role:      "client",
		Password:  string(password),
	}

	database.DB.Create(&user)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

func verifyPassword(password string) error {
	var uppercasePresent bool
	var lowercasePresent bool
	var numberPresent bool
	var specialCharPresent bool
	const minPassLength = 8
	const maxPassLength = 15
	var passLen int
	var errorString string

	for _, ch := range password {
		switch {
		case unicode.IsNumber(ch):
			numberPresent = true
			passLen++
		case unicode.IsUpper(ch):
			uppercasePresent = true
			passLen++
		case unicode.IsLower(ch):
			lowercasePresent = true
			passLen++
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			specialCharPresent = true
			passLen++
		case ch == ' ':
			passLen++
		}
	}
	appendError := func(err string) {
		if len(strings.TrimSpace(errorString)) != 0 {
			errorString += ", " + err
		} else {
			errorString = err
		}
	}
	if !lowercasePresent {
		appendError("Lowercase letter missing")
	}
	if !uppercasePresent {
		appendError("Uppercase letter missing")
	}
	if !numberPresent {
		appendError("At least one numeric character required")
	}
	if !specialCharPresent {
		appendError("Special character missing")
	}
	if !(minPassLength <= passLen && passLen <= maxPassLength) {
		appendError(fmt.Sprintf("Password length must be between %d to %d characters long", minPassLength, maxPassLength))
	}

	if len(errorString) != 0 {
		return fmt.Errorf(errorString)
	}
	return nil
}

type Claims struct {
	Id             uint
	Email          string
	Role           string
	StandardClaims jwt.StandardClaims
}

func (c Claims) Valid() error {
	//TODO implement me
	panic("implement me")
}

func Login(c *fiber.Ctx) error {
	godotenv.Load()
	SecretKey := os.Getenv("SECRETKEY")
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "incorrect password or email",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  "Error",
			"message": "incorrect password or email",
		})
	}

	claims := &Claims{
		Id:    user.Id,
		Email: user.Email,
		Role:  user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 2),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":    tokenString,
	})
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Second),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func VerifyToken(c *fiber.Ctx) error {
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

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Token is empty, unauthenticated",
		})
	}

	email, _ := claims["Email"]

	var user models.User
	if err := database.DB.Find(&user, "email = ?", email).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "Error",
			"message": "There is an error in finding email method",
		})
	}

	return c.Next()
}

func keyFunc(*jwt.Token) (interface{}, error) {
	SecretKey := os.Getenv("SECRETKEY")
	return []byte(SecretKey), nil
}

func Admin(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	tokenArray := strings.Split(token, "Bearer ")
	a := strings.Join(tokenArray, " ")
	to := strings.TrimSpace(a)

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(to, claims, keyFunc)

	if err != nil {
		return c.Status(fiber.StatusGatewayTimeout).JSON(fiber.Map{
			"status":  "error",
			"message": "token expired",
		})
	}

	if claims["Role"] == "admin" {
		return c.Next()
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "unauthorized",
			"message": "you are not an admin, you are not allowed to access this",
		})
	}
}

func Seller(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	tokenArray := strings.Split(token, "Bearer ")
	a := strings.Join(tokenArray, " ")
	to := strings.TrimSpace(a)

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(to, claims, keyFunc)

	if err != nil {
		return c.Status(fiber.StatusGatewayTimeout).JSON(fiber.Map{
			"status":  "error",
			"message": "token expired",
		})
	}

	if claims["Role"] == "seller" || claims["Role"] == "admin" {
		return c.Next()
	} else if claims["Role"] == "client" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "unauthorized",
			"message": "you are not an admin or seller, you are not allowed to access this",
		})
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": "unauthorized",
		})
	}
}

func Buyer(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	tokenArray := strings.Split(token, "Bearer ")
	a := strings.Join(tokenArray, " ")
	to := strings.TrimSpace(a)

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(to, claims, keyFunc)

	if err != nil {
		return c.Status(fiber.StatusGatewayTimeout).JSON(fiber.Map{
			"status":  "error",
			"message": "token expired",
		})
	}

	if (claims["Role"] == "admin") || (claims["Role"] == "seller") || (claims["Role"] == "client") {
		return c.Next()
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "unauthorized",
			"message": "you are not an admin, seller or client, you are not allowed to access this",
		})
	}
}

func User(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	tokenArray := strings.Split(token, "Bearer ")
	a := strings.Join(tokenArray, " ")
	to := strings.TrimSpace(a)

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(to, claims, keyFunc)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	var user models.User
	database.DB.Find(&user, claims["Id"])

	return c.JSON(user)
}
