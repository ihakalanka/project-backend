package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/controllers"
	"main.go/controllers/sellerController"
)

func Dashboard(app *fiber.App) {

	app.Get("/getCount/:id",controllers.Seller, sellercontroller.GetproductCount)
	

}