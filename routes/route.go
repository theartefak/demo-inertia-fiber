package routes

import "github.com/gofiber/fiber/v2"

type Router interface {
	RegisterRoute(app *fiber.App)
}

func RegisterRoute(app *fiber.App) {
	setup(app, NewWeb())
}

func setup(app *fiber.App, router ...Router) {
	for _, r := range router {
		r.RegisterRoute(app)
	}
}
