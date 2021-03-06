package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/controllers/sellerController"
	"main.go/controllers"

)

func Productroute(app *fiber.App) {
	app.Post("/createProducts", controllers.Seller, sellercontroller.Postproduct)
	app.Get("/getProductByid/:id", controllers.Seller, sellercontroller.Getproductid)
	app.Delete("/deleteProduct/:id", controllers.Seller, sellercontroller.Deleteproduct)
	app.Put("/updateProduct/:id", controllers.Seller, sellercontroller.Updateproduct)
	app.Get("/getProductByUserId/:id", controllers.Seller, sellercontroller.GetProductByUserId)
}

