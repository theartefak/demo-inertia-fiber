package controllers

import (
	"github.com/gofiber/fiber/v2"
	middleware "github.com/theartefak/artefak/app/Middleware"
	models "github.com/theartefak/artefak/app/Models"
	"github.com/theartefak/artefak/database"
)

// Welcome handles the welcome page request.
func Welcome(c *fiber.Ctx) error {
	// Retrieve all users from the database
	users := []models.User{}
	database.DB.Find(&users)

	// Render the welcome page with greeting and user data
	return c.Render("Index", fiber.Map{
		"greeting" : "Hello World",
		"users"    : users,
	})
}

// CreateDummyUser creates a new dummy user based on the request.
func CreateDummyUser(c *fiber.Ctx) error {
	// Create a new user instance
	user := new(models.User)

	// Validate the request data using the middleware
	validate, parser := middleware.Validate("Index", c, user)

	// If parser errors exist, return a Internal server Error response
	if parser != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(parser)
	}

	// If validation errors exist, return a Status Found response
	if validate != nil {
		return c.Status(fiber.StatusOK).JSON(validate)
	}

	// Save the user to the database
	database.DB.Create(&user)

	// Return a success response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Success"})
}
