package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/theartefak/artefak/database"
	"github.com/theartefak/artefak/routes"
	"github.com/theartefak/inertia-fiber"
)

type GlobalErrorHandlerResp struct {
	Success bool `json:"success"`
	Message string `json:"message"`
}

func Run() *fiber.App {
	database.InitDB()

	engine := inertia.New()

	artefak := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusBadRequest).JSON(GlobalErrorHandlerResp{
				Success: false,
				Message: err.Error(),
			})
		},
		Views: engine,
	})

	artefak.Use(favicon.New(favicon.Config{
		File: "./public/favicon.ico",
	}))

	artefak.Use(cors.New())
	artefak.Use(csrf.New())
	artefak.Use(engine.Middleware())
	artefak.Use(helmet.New(helmet.Config{
		CrossOriginOpenerPolicy: "cross-origin",
		CrossOriginResourcePolicy: "cross-origin",
		OriginAgentCluster: "?0",
	}))
	artefak.Use(logger.New())
	artefak.Static("/assets", "public/build/assets").Name("asset")

	routes.RegisterRoute(artefak)

	return artefak
}
