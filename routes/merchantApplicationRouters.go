package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/controllers/merchantapplicationcontroller"
	"main.go/controllers"
)

func MerchantApplicationroute(app *fiber.App){
	app.Use(controllers.VerifyToken)
	app.Get("/getMerchantApplication",controllers.Seller,merchantapplicationcontroller.GetMerchantApplication)
	 
}