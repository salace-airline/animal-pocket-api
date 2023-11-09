package handlers

import (
	"github.com/salace-airline/animalpocketresources/cmd/database"
	"github.com/salace-airline/animalpocketresources/cmd/models"

	"github.com/gofiber/fiber/v2"
)

// @Summary      Get bugs
// @Description  Get all bugs which exists
// @Tags         resources
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.BugSwagger
// @Failure      422  {string}  string    "error"
// @Failure      500  {string}  string    "error"
// @Router       /bug [get]
func GetBugs(c *fiber.Ctx) error {
	bugs := []models.Bug{}
	database.DB.Db.Find(&bugs)

	return c.Status(200).JSON(bugs)
}
