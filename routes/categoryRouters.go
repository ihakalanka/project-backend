package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/controllers/admincontrollers"
)

func Categoryroute(app *fiber.App) {
	app.Get("/getCategory", admincontrollers.Getcat)
	app.Post("/createCategory", admincontrollers.Postcat)
	app.Get("/getCategoryByid/:id", admincontrollers.Getcatid)
	app.Delete("/deleteCategory/:id", admincontrollers.Deletecat)
	app.Put("updateCategory/:id", admincontrollers.Updatecat)
}
