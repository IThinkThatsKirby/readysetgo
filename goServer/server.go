package main

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Start your fiber server
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Howdy, World!")
	})
	user := app.Group("/:user")

	user.Get("/todo/", func(c *fiber.Ctx) error {
		return c.SendString("Your todos for " + c.Params("user"))
	})
	user.Get("/todo/:id", func(c *fiber.Ctx) error {
		return c.SendString("a specifc task with data")
	})
	user.Get("/todone/", func(c *fiber.Ctx) error {
		return c.SendString("Your archived and completed tasks")
	})
	user.Get("/todone/:id", func(c *fiber.Ctx) error {
		return c.SendString("Data on the completed task")
	})
	app.Listen(":8080")
}
