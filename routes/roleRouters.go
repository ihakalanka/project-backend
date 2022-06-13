package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/controllers/admincontrollers"
	"main.go/controllers/sellercontroller"
)
func Roleroute(app *fiber.App){
	app.Get("/getRole", admincontrollers.Getrole)
	app.Post("/createRole",admincontrollers.Postrole)
	app.Get("/getRoleId/:id",admincontrollers.Getroleid)
	app.Delete("/deleteRole/:id", admincontrollers.Deleterole)
	app.Put("updateRole/:id", admincontrollers.Updaterole)
	app.Get("/getUser",sellercontroller.GetAllSellers)
	app.Delete("deleteUser/:id",sellercontroller.DeleteSeller)
	app.Put("/updateUser/:id",sellercontroller.UpdateSeller)
}