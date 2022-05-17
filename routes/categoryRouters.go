package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/controllers/admincontrollers"
)

func Categoryroute(app *fiber.App) {
	app.Get("/getCategory",admincontrollers.Get)
	app.Post("/createCategory",admincontrollers.Post)
}
