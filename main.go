package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message":   "My name is Terry",
			"timestamp": time.Now().Unix(),
		})
	})

	app.Listen(":3000")
}
