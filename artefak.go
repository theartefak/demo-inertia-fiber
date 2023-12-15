package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/theartefak/inertia-fiber"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(csrf.New())
	app.Use(helmet.New())
	app.Use(limiter.New())
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(inertia.Inertia())
	})

    app.Listen("127.0.1.1:8000")
}
