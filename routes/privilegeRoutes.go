package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/controllers"
	"main.go/controllers/privilegecontroller"
)

func Privilegeroute(app *fiber.App) {

	app.Get("/getPrivilege", controllers.Admin, privilegecontroller.GetPrivilege)
	app.Post("/createPrivilege", controllers.Admin, privilegecontroller.PostPrivilege)
	app.Get("/getPrivilegeByid/:id", controllers.Admin, privilegecontroller.GetPrivilegeid)
	app.Delete("/deletePrivilege/:id", controllers.Admin, privilegecontroller.DeletePrivilege)
	app.Put("/updatePrivilege/:id", controllers.Admin, privilegecontroller.UpdatePrivilege)
}
