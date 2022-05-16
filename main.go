package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"main.go/database"
	"main.go/routes"
)

func main() {
	database.Connect()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
		
	}))
	routes.Route(app)
	routes.Categoryroute(app)
	log.Fatal(app.Listen(":8080"))
}
