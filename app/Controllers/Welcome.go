package controllers

import "github.com/gofiber/fiber/v2"

func Welcome(c *fiber.Ctx) error {
	return c.Render("Index", fiber.Map{
		"greeting": "Hello World",
	})
}
