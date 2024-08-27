package handlers

import (
	"github.com/salace-airline/animalpocketapi/database"
	"github.com/salace-airline/animalpocketapi/models"

	"github.com/gofiber/fiber/v2"
)

func GetFishes(c *fiber.Ctx) error {
	fishes := []models.Fish{}
	database.DB.Db.Find(&fishes)

	return c.Status(fiber.StatusOK).JSON(fishes)
}
