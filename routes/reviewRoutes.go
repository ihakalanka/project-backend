package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/controllers"
	"main.go/controllers/customerController"
)

func ReviewRoutes(app *fiber.App) {
	app.Use(controllers.VerifyToken)

	app.Post("/api/createreview", customerController.CreateReview)
}
