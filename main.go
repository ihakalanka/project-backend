package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"main.go/database"
	"main.go/routes"
)

func main() {
	database.Connect()
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
	}))
	routes.Route(app)
	routes.Categoryroute(app)
	routes.Productroute(app)
	routes.Roleroute(app)
	routes.Merchantroute(app)

	log.Fatal(app.Listen(":8080"))
}
