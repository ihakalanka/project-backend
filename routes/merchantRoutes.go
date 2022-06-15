package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/controllers/merchantcontroller"
	 
)

func Merchantroute(app *fiber.App){
	app.Get("/getMerchant",merchantcontroller.GetMerchant)
	app.Post("/createMerchant",merchantcontroller.PostMerchant)
	app.Get("/getMerchantByid/:id",merchantcontroller.GetMerchantid)
	app.Put("/updateMerchant/:id",merchantcontroller.UpdateMerchant)
}