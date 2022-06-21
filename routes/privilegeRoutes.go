package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/controllers/privilegecontroller"
	"main.go/controllers"
)

func Privilegeroute(app *fiber.App){

	app.Use(controllers.VerifyToken)
	app.Get("/getPrivilege",controllers.Admin,privilegecontroller.GetPrivilege)
	app.Post("/createPrivilege",controllers.Admin,privilegecontroller.PostPrivilege)
	app.Get("/getPrivilegeByid/:id",controllers.Admin,privilegecontroller.GetPrivilegeid)
	app.Delete("/deletePrivilege/:id",controllers.Admin,privilegecontroller.DeletePrivilege)
	app.Put("/updatePrivilege/:id",controllers.Admin,privilegecontroller.UpdatePrivilege)
}