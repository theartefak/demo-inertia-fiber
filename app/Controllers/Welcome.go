package controllers

import (
	"github.com/gofiber/fiber/v2"
	models "github.com/theartefak/artefak/app/Models"
	"github.com/theartefak/artefak/database"
)

func Welcome(c *fiber.Ctx) error {
	users := []models.User{}
	database.DB.Find(&users)

	return c.Render("Index", fiber.Map{
		"greeting": "Hello World",
		"users": users,
	})
}

func CreateDummyUser(c *fiber.Ctx) error {
	user := models.User{
		Name: "Test User",
		Email: "test@user.com",
		Password: "12345678",
	}

	database.DB.Create(&user)

	return c.SendString("User Successfully created.")
}
