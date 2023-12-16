package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/theartefak/inertia-fiber"
)

func main() {
	engine := inertia.New()

	artefak := fiber.New(fiber.Config{
		Views: engine,
	})

	artefak.Use(favicon.New(favicon.Config{
		File: "./public/favicon.ico",
	    URL: "/favicon.ico",
	}))

	artefak.Use(cors.New())
	artefak.Use(csrf.New())
	artefak.Use(helmet.New())
	artefak.Use(limiter.New())
	artefak.Use(logger.New())
	artefak.Use(engine.Middleware())
	artefak.Static("/assets", "public/build/assets")

	artefak.Get("/", func(c *fiber.Ctx) error {
		return c.Render("Index", fiber.Map{
			"greeting": "Hello World",
		})
	}).Name("Welcome")

    artefak.Listen("127.0.1.1:8000")
}
