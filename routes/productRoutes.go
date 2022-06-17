package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/controllers"
	"main.go/controllers/sellercontroller"
)

func Productroute(app *fiber.App) {
	app.Use(controllers.VerifyToken)


	app.Get("/getProducts", controllers.Seller, sellercontroller.Getproduct)
	app.Post("/createProducts", controllers.Seller, sellercontroller.Postproduct)
	app.Get("/getProductByid/:id", controllers.Seller, sellercontroller.Getproductid)
	app.Delete("/deleteProduct/:id", controllers.Seller, sellercontroller.Deleteproduct)
	app.Put("/updateProduct", controllers.Seller, sellercontroller.Updateproduct)
}
