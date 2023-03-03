package main

import (
	"github.com/salace-airline/ressources/handlers"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.GetWelcome)
	app.Get("/fishes", handlers.GetFishes)
	app.Post("/fish", handlers.CreateFish)
}
