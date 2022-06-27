package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/controllers/customerController"
)

func ReviewRoutes(app *fiber.App) {
	app.Post("/api/createReview", customerController.CreateReview)
	app.Delete("/api/deleteReviewByItem/:id", customerController.DeleteReview)
}
