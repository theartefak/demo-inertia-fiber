package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/theartefak/inertia-fiber"
)

func Run() *fiber.App {
	engine := inertia.New()

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(favicon.New(favicon.Config{
		File: "./public/favicon.ico",
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	app.Use(csrf.New())
	app.Use(engine.Middleware())
	app.Use(helmet.New())
	app.Use(logger.New())
	app.Static("/assets", "public/build/assets")

	return app
}
