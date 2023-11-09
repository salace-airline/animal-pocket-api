package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// @Summary Welcome to you!
// @Description Because you need a sweet start before using our API.
// @Tags basic
// @Accept */*
// @Consumes json
// @Produces json
// @Success 200 {string} string "Welcome to Animal Pocket API!"
// @Router / [get]
func GetWelcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to Animal Pocket API!")
}
