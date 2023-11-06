package handlers

import (
	_ "github.com/salace-airline/animalpocketresources/cmd/docs"

	"github.com/gofiber/fiber/v2"
)

// @Summary Welcome to you!
// @Description Because you need a sweet start before using our API.
// @Tags root
// @Accept */*
// @Consumes application/json
// @Produces application/json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func GetWelcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to Animal Pocket API!")
}
