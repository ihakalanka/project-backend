package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/controllers"
)

func UserRoutes(app *fiber.App) {
	app.Get("/api/user", controllers.User)
}
