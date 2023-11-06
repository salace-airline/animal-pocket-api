package main

import (
	_ "github.com/salace-airline/animalpocketresources/cmd/docs"
	"github.com/salace-airline/animalpocketresources/handlers"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	// Welcome!
	app.Get("/", handlers.GetWelcome)

	// Authentication
	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)
	app.Post("/logout", handlers.Logout)

	// User
	app.Get("/user", handlers.GetActualUser)
	/* New api routes coming soon !
	app.Patch("/user/:id", handlers.UpdateUser)
	app.Delete("/user/:id", handlers.DeleteUser)*/

	// Ressources
	app.Get("/fish", handlers.GetFishes)
	app.Get("/bug", handlers.GetBugs)
	app.Get("/sea-creature", handlers.GetSeaCreatures)
}
