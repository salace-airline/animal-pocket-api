package main

import (
	"github.com/salace-airline/animalpocketresources/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// @title Animal Pocket API
	// @version 1.0
	// @description This is the API used by the iOS application named Animal Pocket.
	
	// Connect and fill database with resources
	database.ConnectDb()
	database.Migrate()
	database.IntegrateResources()

	app := fiber.New()

	// Allow cors for cookie
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// Initialize routes and run
	setupRoutes(app)

	app.Listen(":3000")
}
