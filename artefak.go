package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	models "github.com/theartefak/artefak/app/Models"
	"github.com/theartefak/artefak/database"
	"github.com/theartefak/artefak/routes"
	"github.com/theartefak/inertia-fiber"
)

type App struct {
	*fiber.App

	DB      *database.Database
	Inertia *inertia.Engine
}

func main() {
	inertia := inertia.New()

	app := App{
		App: fiber.New(fiber.Config{
			Views: inertia,
		}),
	}

	// Initialize database
	db, err := database.New()

	// Auto-migrate database models
	if err != nil {
		fmt.Println("failed to connect to database:", err.Error())
	} else {
		if db == nil {
			fmt.Println("failed to connect to database: db variable is nil")
		} else {
			app.DB = db
			err = app.DB.AutoMigrate(&models.User{})
			if err != nil {
				fmt.Println("failed to automigrate user model:", err.Error())
				return
			}
		}
	}

	app.Use(favicon.New(favicon.Config{
		File: "./public/favicon.ico",
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))
	app.Use(csrf.New(csrf.Config{
		KeyLookup: "header:X-XSRF-TOKEN",
		CookieName: "XSRF-TOKEN",
		SingleUseToken: true,
	}))
	app.Use(helmet.New(helmet.Config{
		ReferrerPolicy: "strict-origin-when-cross-origin",
	}))
	app.Use(inertia.Middleware())
	app.Use(logger.New())
	app.Use(recover.New())
	app.Static("/assets", "public/build/assets").Name("asset")

	web := app.Group("")
	routes.RegisterRoute(web, inertia, app.DB)

	// Close any connections on interrupt signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		app.exit()
	}()

	// Start listening on the specified address
	err = app.Listen("127.0.0.1:8000")
	if err != nil {
		app.exit()
	}
}

// Stop the Fiber application
func (app *App) exit() {
	_ = app.Shutdown()
}
