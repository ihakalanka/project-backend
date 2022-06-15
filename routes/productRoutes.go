package routes

import (
	"github.com/gofiber/fiber/v2"
	// "main.go/controllers"
	"main.go/controllers/sellercontroller"
)

func Productroute(app *fiber.App) {
	// app.Use(controllers.VerifyToken)

	app.Get("/getProducts", sellercontroller.Getproduct)
	app.Post("/createProducts", sellercontroller.Postproduct)
	app.Get("/getProductByid/:id", sellercontroller.Getproductid)
	app.Delete("/deleteProduct/:id", sellercontroller.Deleteproduct)
	app.Put("/updateProduct/:id", sellercontroller.Updateproduct)
}
