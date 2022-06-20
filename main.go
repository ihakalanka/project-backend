package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"main.go/database"
	"main.go/routes"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	database.Connect()

	godotenv.Load()

	app := fiber.New()
	 
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Route(app)
	routes.Categoryroute(app)
	routes.Productroute(app)
	routes.Roleroute(app)
	routes.Merchantroute(app)

	port := os.Getenv("PORT")
	log.Fatal(app.Listen(port))
}
	