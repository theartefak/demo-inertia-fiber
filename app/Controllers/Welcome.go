package controllers

import (
	"github.com/gofiber/fiber/v2"
	middleware "github.com/theartefak/artefak/app/Middleware"
	models "github.com/theartefak/artefak/app/Models"
	"github.com/theartefak/artefak/database"
)

// Welcome handles the welcome page request.
func Welcome(db *database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Retrieve all users from the database
		users := []models.User{}
		db.Find(&users)

		// Render the welcome page with greeting and user data
		return c.Render("Index", fiber.Map{
			"greeting" : "Hello World",
			"users"    : users,
		})
	}
}

// CreateDummyUser creates a new dummy user based on the request.
func CreateDummyUser(db *database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Create a new user instance
		user := new(models.User)

		// If validation errors exist, return a Status Found response
		if validate, err := middleware.Validate("Index", user)(c); err != nil {
			return c.Status(302).JSON(validate)
		}

		// Save the user to the database
		db.Create(&user)

		return nil
	}
}
