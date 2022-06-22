package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"log"
	"main.go/controllers"
	"main.go/database"
	"main.go/routes"
	"os"
)

func main() {
	database.Connect()

	godotenv.Load()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
	}))

	routes.Route(app)
	routes.ViewProdRoutes(app)
	app.Use(controllers.VerifyToken)
	routes.UserRoutes(app)
	routes.Categoryroute(app)
	routes.Productroute(app)
	routes.Roleroute(app)
	routes.Cartroute(app)
	routes.WishListroute(app)
	routes.Dashboard(app)
	routes.ReviewRoutes(app)

	port := os.Getenv("PORT")
	log.Fatal(app.Listen(port))
}
