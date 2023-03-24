package main

import (
	"github.com/salace-airline/animalpocketresources/handlers"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.GetWelcome)
	app.Get("/fishes", handlers.GetFishes)
	app.Get("/bugs", handlers.GetBugs)
	app.Get("/sea-creatures", handlers.GetSeaCreatures)
}
