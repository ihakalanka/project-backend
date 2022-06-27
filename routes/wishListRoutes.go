package routes

import (
	"github.com/gofiber/fiber/v2"
	// "main.go/controllers"
	"main.go/controllers/customerController"
)

func WishListroute(app *fiber.App) {
	// app.Use(controllers.VerifyToken)

	
	app.Post("/createList", customerController.Postlist)
	
	app.Delete("/deleteList/:id",customerController.Deletelist)
	app.Get("/getListbyUserId/:id", customerController.GetlistbyUserId)
}