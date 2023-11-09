package handlers

import (
	"github.com/salace-airline/animalpocketresources/cmd/database"
	"github.com/salace-airline/animalpocketresources/cmd/models"

	"github.com/gofiber/fiber/v2"
)

func GetSeaCreatures(c *fiber.Ctx) error {
	seaCreatures := []models.SeaCreature{}
	database.DB.Db.Find(&seaCreatures)

	return c.Status(200).JSON(seaCreatures)
}
