package handlers

import (
	"github.com/salace-airline/animalpocketresources/cmd/database"
	"github.com/salace-airline/animalpocketresources/cmd/models"

	"github.com/gofiber/fiber/v2"
)

func GetFishes(c *fiber.Ctx) error {
	fishes := []models.Fish{}
	database.DB.Db.Find(&fishes)

	return c.Status(200).JSON(fishes)
}
