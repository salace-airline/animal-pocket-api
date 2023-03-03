package handlers

import (
	"github.com/salace-airline/ressources/database"
	"github.com/salace-airline/ressources/models"

	"github.com/gofiber/fiber/v2"
)

func GetFishes(c *fiber.Ctx) error {
	fishes := []models.Fish{}
	database.DB.Db.Find(&fishes)

	return c.Status(200).JSON(fishes)
}

func CreateFish(c *fiber.Ctx) error {
	fish := new(models.Fish)
	if err := c.BodyParser(fish); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fish)

	return c.Status(200).JSON(fish)
}
