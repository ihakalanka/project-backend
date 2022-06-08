package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/controllers"
	"main.go/controllers/sellercontroller"
)

<<<<<<< HEAD
func Productroute(app *fiber.App){
	app.Get("/getProducts",sellercontroller.Getproduct)
	app.Post("/createProducts",sellercontroller.Postproduct)
	app.Get("/getProductByid/:id",sellercontroller.Getproductid)
	app.Delete("/deleteProduct/:id",sellercontroller.Deleteproduct)
	app.Put("/updateProduct/:id",sellercontroller.Updateproduct)
=======
func Productroute(app *fiber.App) {
	app.Use(controllers.VerifyToken)

	app.Get("/api/getProducts", sellercontroller.Getproduct)
	app.Post("/createProducts", sellercontroller.Postproduct)
	app.Get("/getProductByid/:id", sellercontroller.Getproductid)
	app.Delete("/deleteProduct/:id", sellercontroller.Deleteproduct)
	app.Put("/updateProduct", sellercontroller.Updateproduct)
>>>>>>> 6024d996114591d2144db250a6e8f560980d3518
}
