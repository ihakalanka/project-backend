package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/controllers"
	"main.go/controllers/merchantcontroller"
)

func Merchantroute(app *fiber.App) {
	app.Get("/getMerchant", controllers.Seller, merchantcontroller.GetMerchant)
	app.Post("/createMerchant", controllers.Seller, merchantcontroller.PostMerchant)
	app.Get("/getMerchantByid/:id", controllers.Seller, merchantcontroller.GetMerchantByUserid)
	app.Put("/updateMerchant/:id", controllers.Seller, merchantcontroller.UpdateMerchant)
	app.Get("/getDataById/:id", controllers.Seller, merchantcontroller.GetMerchantid)
}
