package main

import (
	"github.com/salace-airline/animalpocketapi/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Connect and fill database with resources
	database.ConnectDb()
	database.Migrate()
	database.IntegrateResources()

	app := fiber.New()

	// Allow cors for cookie
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: true,
	}))

	// Initialize routes and run
	setupRoutes(app)

	app.Listen(":3000")
}
