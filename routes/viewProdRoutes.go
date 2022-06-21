package routes

import (
	"github.com/gofiber/fiber/v2"
	sellercontroller "main.go/controllers/sellerController"
)

func ViewProdRoutes(app *fiber.App) {
	app.Get("/getProducts", sellercontroller.Getproduct)
	app.Get("/getProductByid/:id", sellercontroller.Getproductid)
}
