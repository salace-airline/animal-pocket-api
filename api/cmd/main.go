package main

import (
	_ "github.com/salace-airline/animalpocketresources/cmd/docs"
	"github.com/salace-airline/animalpocketresources/database"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// @title Animal Pocket API
	// @version 1.0
	// @description You can found all the ressources in Animal Crossing: New Horizons. It is used by the iOS application named Animal Pocket.
	// @host 13.39.189.214:3000
	// @BasePath /
	// @schemes http

	// Connect and fill database with resources
	database.ConnectDb()
	database.Migrate()
	database.IntegrateResources()

	app := fiber.New()

	app.Get("/swagger/*", swagger.New(swagger.Config{
		BasePath: "/",
		FilePath: "./cmd/docs/swagger.json",
		Path:     "docs",
		Title:    "Swagger API Docs",
	}))

	// Allow cors for cookie
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// Initialize routes and run
	setupRoutes(app)

	app.Listen(":3000")
}
