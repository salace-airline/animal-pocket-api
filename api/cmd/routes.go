package main

import (
	"github.com/salace-airline/animalpocketresources/handlers"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.GetWelcome)
	app.Get("/fish", handlers.GetFishes)
	app.Get("/bug", handlers.GetBugs)
	app.Get("/sea-creature", handlers.GetSeaCreatures)
}
