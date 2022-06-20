package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/controllers"
	"main.go/controllers/customerController"
)

func ReviewRoutes(app *fiber.App) {
	app.Use(controllers.VerifyToken)

	app.Post("/api/createReview", customerController.CreateReview)
	app.Get("/api/getAverageRating/:id", customerController.ViewAverageRating)
	app.Get("/api/getAllReviewsByItem/:id", customerController.ViewAllReviewsByItem)
}
