package main

import "github.com/gofiber/fiber/v2"

func main() {
	// Start your fiber server
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Howdy, World!")
	})
	app.Listen(":8080")
}
