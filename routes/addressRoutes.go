package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/controllers"
	"main.go/controllers/addressController"
)

func Addressroute(app *fiber.App) {
	app.Get("/getAddress", controllers.Seller, addressControllers.Getcat)
	app.Post("/setAddress", controllers.Seller, addressControllers.Postcat)

}
