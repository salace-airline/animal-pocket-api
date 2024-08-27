package main

import (
	"github.com/salace-airline/animalpocketapi/handlers"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.GetWelcome)

	// Authentication
	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)
	app.Post("/logout", handlers.Logout)

	// User
	app.Get("/user", handlers.GetActualUser)
	app.Patch("/user", handlers.UpdateUser)
	app.Patch("/user/", handlers.UpdateUser)
	app.Patch("/user/fish", handlers.UpdateUserFish)
	app.Patch("/user/bug", handlers.UpdateUserBug)
	app.Patch("/user/sea-creature", handlers.UpdateUserSeaCreature)
	/* New api routes coming soon !
	app.Delete("/user/:id", handlers.DeleteUser)*/

	// Ressources
	app.Get("/fish", handlers.GetFishes)
	app.Get("/bug", handlers.GetBugs)
	app.Get("/sea-creature", handlers.GetSeaCreatures)
}
