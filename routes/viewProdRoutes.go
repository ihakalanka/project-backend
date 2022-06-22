package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/controllers/customerController"
	sellercontroller "main.go/controllers/sellerController"
)

func ViewProdRoutes(app *fiber.App) {
	app.Get("/getProducts", sellercontroller.Getproduct)
	app.Get("/getProductByid/:id", sellercontroller.Getproductid)

	app.Get("/api/getAverageRating/:id", customerController.ViewAverageRating)
	app.Get("/api/getAllReviewsByItem/:id", customerController.ViewAllReviewsByItem)
}
