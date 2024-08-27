package handlers

import (
	"github.com/salace-airline/animalpocketapi/database"
	"github.com/salace-airline/animalpocketapi/models"

	"github.com/gofiber/fiber/v2"
)

func GetSeaCreatures(c *fiber.Ctx) error {
	seaCreatures := []models.SeaCreature{}
	database.DB.Db.Find(&seaCreatures)

	return c.Status(fiber.StatusOK).JSON(seaCreatures)
}
