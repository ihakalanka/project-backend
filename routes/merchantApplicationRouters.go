package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/controllers/merchantapplicationcontroller"
	 
)

func MerchantApplicationroute(app *fiber.App){
	app.Get("/getMerchant",merchantapplicationcontroller.GetMerchantApplication)
	 
}