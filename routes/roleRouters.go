package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/controllers"
	"main.go/controllers/admincontrollers"
)

func Roleroute(app *fiber.App) {
	app.Use(controllers.VerifyToken)

	app.Get("/getRole", admincontrollers.Getrole)
	app.Post("/createRole", admincontrollers.Postrole)
	app.Get("/getRoleId/:id", admincontrollers.Getroleid)
	app.Delete("/deleteRole/:id", admincontrollers.Deleterole)
	app.Put("updateRole/:id", admincontrollers.Updaterole)
}
