package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/controllers"
	"main.go/controllers/admincontrollers"
)

func Categoryroute(app *fiber.App) {
	app.Use(controllers.VerifyToken)

	app.Get("/getCategory", controllers.Admin, admincontrollers.Getcat)
	app.Post("/createCategory", controllers.Admin, admincontrollers.Postcat)
	app.Get("/getCategoryByid/:id", controllers.Admin, admincontrollers.Getcatid)
	app.Delete("/deleteCategory/:id", controllers.Admin, admincontrollers.Deletecat)
	app.Put("updateCategory/:id", controllers.Admin, admincontrollers.Updatecat)
}
