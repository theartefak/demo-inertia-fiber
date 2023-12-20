package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/theartefak/artefak/bootstrap"
)

func main() {
	app := bootstrap.Run()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("Index", fiber.Map{
			"greeting": "Hello World",
		})
	}).Name("Welcome")

    app.Listen("127.0.1.1:8000")
}
