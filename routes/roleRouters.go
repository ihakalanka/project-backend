package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/controllers"
	"main.go/controllers/admincontrollers"
)

func Roleroute(app *fiber.App) {
	app.Get("/getRole", controllers.Admin, admincontrollers.Getrole)
	app.Post("/createRole", controllers.Admin, admincontrollers.Postrole)
	app.Get("/getRoleId/:id", controllers.Admin, admincontrollers.Getroleid)
	app.Delete("/deleteRole/:id", controllers.Admin, admincontrollers.Deleterole)
	app.Put("updateRole/:id", controllers.Admin, admincontrollers.Updaterole)
}
