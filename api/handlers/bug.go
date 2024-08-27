package handlers

import (
	"github.com/salace-airline/animalpocketapi/database"
	"github.com/salace-airline/animalpocketapi/models"

	"github.com/gofiber/fiber/v2"
)

func GetBugs(c *fiber.Ctx) error {
	bugs := []models.Bug{}
	database.DB.Db.Find(&bugs)

	return c.Status(fiber.StatusOK).JSON(bugs)
}
