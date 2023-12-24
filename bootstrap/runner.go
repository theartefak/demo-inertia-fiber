package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/theartefak/artefak/database"
	"github.com/theartefak/artefak/routes"
	"github.com/theartefak/inertia-fiber"
)

func Run() *fiber.App {
	database.InitDB()

	engine := inertia.New()

	artefak := fiber.New(fiber.Config{
		Views: engine,
	})

	artefak.Use(favicon.New(favicon.Config{
		File: "./public/favicon.ico",
	}))


	artefak.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))
	artefak.Use(csrf.New(csrf.Config{
		KeyLookup: "header:X-XSRF-TOKEN",
		CookieName: "XSRF-TOKEN",
		SingleUseToken: true,
	}))
	artefak.Use(helmet.New(helmet.Config{
		ReferrerPolicy: "strict-origin-when-cross-origin",
	}))
	artefak.Use(engine.Middleware())
	artefak.Use(logger.New())
	artefak.Use(recover.New())
	artefak.Static("/assets", "public/build/assets").Name("asset")

	routes.RegisterRoute(artefak)

	return artefak
}
