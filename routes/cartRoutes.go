package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/controllers"
	"main.go/controllers/customerController"
)

func Cartroute(app *fiber.App) {
	app.Use(controllers.VerifyToken)

	
	app.Post("/createCart", customerController.Postcart)
	app.Put("/updateCart/:id", customerController.Updatecart)
	app.Delete("/deleteCart/:id",customerController.Deletecart)
	app.Get("/getCartByUserId/:id", customerController.Getcart)
}