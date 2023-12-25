package routes

import (
	"github.com/gofiber/fiber/v2"
	controllers "github.com/theartefak/artefak/app/Controllers"
	"github.com/theartefak/artefak/database"
	"github.com/theartefak/inertia-fiber"
)

func RegisterRoute(route fiber.Router, inertia *inertia.Engine, db *database.Database) {
	route.Get("", controllers.Welcome(db)).Name("welcome")
	route.Post("create-dummy-user", controllers.CreateDummyUser(db)).Name("create.dummy.user")
}
