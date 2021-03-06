package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/controllers"
	"main.go/controllers/admincontrollers"
)

func Categoryroute(app *fiber.App) {
	app.Get("/getCategory", controllers.Seller, admincontrollers.Getcat)
	app.Post("/createCategory", controllers.Admin, admincontrollers.Postcat)
	app.Get("/getCategoryByid/:id", controllers.Admin, admincontrollers.Getcatid)
	app.Put("updateCategory/:id", controllers.Admin, admincontrollers.Updatecat)
	app.Delete("/deleteCategory/:id", controllers.Admin, admincontrollers.Deletecat)

}
