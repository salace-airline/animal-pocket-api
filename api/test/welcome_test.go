package test

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetWelcome(t *testing.T) {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("Welcome to Animal Pocket API!")
	})

	r := httptest.NewRequest("GET", "/", nil)
	resp, _ := app.Test(r, -1)
	assert.Equal(t, 200, resp.StatusCode)
}
