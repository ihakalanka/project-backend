package routes

import (
	"github.com/gofiber/fiber/v2"
    "main.go/controllers"
    "main.go/controllers/companycontrollers"
)

func CompanyRoutes(router *fiber.App) {

	app.Use(controllers.VerifyToken)

	app.Get("/getCompany",  companycontrollers.Getcat)
    // app.Post("/createCompany", controllers.Admin, companycontrollers.Postcat)
    // app.Get("/getCompanyByid/:id", controllers.Admin, companycontrollers.Getcatid)
    // app.Put("updateCompany/:id",controllers.Admin, companycontrollers.Updatecat)
    // app.Delete("/deleteCompany/:id", controllers.Admin, companycontrollers.Deletecat)
    

}
