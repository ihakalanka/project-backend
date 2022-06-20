package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/controllers/privilegecontroller"
	 
)

func Privilegeroute(app *fiber.App){
	app.Get("/getPrivilege",privilegecontroller.GetPrivilege)
	app.Post("/createPrivilege",privilegecontroller.PostPrivilege)
	app.Get("/getPrivilegeByid/:id",privilegecontroller.GetPrivilegeid)
	app.Delete("/deletePrivilege/:id",privilegecontroller.DeletePrivilege)
	app.Put("/updatePrivilege/:id",privilegecontroller.UpdatePrivilege)
}