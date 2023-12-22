package routes

import (
	"github.com/gofiber/fiber/v2"
	controllers "github.com/theartefak/artefak/app/Controllers"
)

type Web struct{}

func (w *Web) RegisterRoute(app *fiber.App) {
	route := app.Group("")

	route.Get("", controllers.Welcome).Name("welcome")
	route.Post("create-dummy-user", controllers.CreateDummyUser).Name("create.dummy.user")
}

func NewWeb() *Web {
	return &Web{}
}
