package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/controllers"
	_ "main.go/controllers"
	companyControllers "main.go/controllers/companyConroller"
)

func CompanyRoutes(app *fiber.App) {
	app.Get("/getCompany", controllers.Seller, companyControllers.Getcat)
	app.Post("/createCompany", controllers.Seller, companyControllers.Postcat)
	app.Get("/getCompanyByid/:id", controllers.Seller, companyControllers.Getcatid)
	app.Delete("/deleteCompany/:id", controllers.Seller, companyControllers.Deletecat)
}
