package main

import (
	"github.com/salace-airline/animalpocketresources/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	database.Migrate()
	database.IntegrateResources()

	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
