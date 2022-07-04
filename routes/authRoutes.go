package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/controllers"
)

func Route(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Post("/api/logout", controllers.Logout)
	app.Post("/api/forgot", controllers.Forgot)
	app.Post("/api/reset", controllers.Reset)

	app.Post("/api/verify", controllers.VerifyMail)

	app.Post("/api/companyRegister", controllers.CompanyRegister)

	app.Post("/api/sendEmail", controllers.SendEmail)
}
