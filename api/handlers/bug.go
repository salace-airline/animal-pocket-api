package handlers

import (
	"github.com/salace-airline/animalpocketresources/database"
	"github.com/salace-airline/animalpocketresources/models"

	"github.com/gofiber/fiber/v2"
)

func GetBugs(c *fiber.Ctx) error {
	bugs := []models.Bug{}
	database.DB.Db.Find(&bugs)

	return c.Status(200).JSON(bugs)
}
