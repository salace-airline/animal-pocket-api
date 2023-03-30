package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetWelcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to Animal Pocket API!")
}
