package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/controllers"
)

func SearchRoutes(app *fiber.App){
	app.Post("/searchItem", controllers.GetAllSearchResult)
}