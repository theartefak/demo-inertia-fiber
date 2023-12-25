package middleware

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Constants for validation error messages
const (
	emailTag    = "must be a valid email address"
	minTag      = "must be at least %s characters"
	requiredTag = "field is required"
)

// ErrorValidation represents a single validation error.
type ErrorValidation struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

var validate = validator.New()

// Validate is a middleware function for request body validation using the provided model.
// It returns a validation error response if there are errors in the request body.
func Validate(component string, model interface{}) func(c *fiber.Ctx) (fiber.Map, error) {
	return func(c *fiber.Ctx) (fiber.Map, error) {
		response := fiber.Map{
			"component" : component,
			"props"     : fiber.Map{"errors": fiber.Map{}},
			"url"       : strings.TrimPrefix(c.Get("Referer"), c.BaseURL()),
			"version"   : c.GetRespHeader("X-Inertia-Version"),
		}

		// Parse request body into the provided model
		if err := c.BodyParser(model); err != nil {
			response["props"].(fiber.Map)["errors"].(fiber.Map)["message"] = "Failed to parse request body"
			return response, err
		}

		// Validate the parsed model and return any validation errors

		if validationErrors, err := validateStruct(validate, model); len(validationErrors) > 0 && err != nil {
			// Map validation errors to user-friendly error messages
			for _, err := range validationErrors {
				field := strings.ToLower(err.Field)

				var tag = err.Tag
				switch tag {
				case "email":
					tag = emailTag
				case "min":
					tag = fmt.Sprintf(minTag, err.Param)
				default:
					tag = requiredTag
				}

				response["props"].(fiber.Map)["errors"].(fiber.Map)[field] = fmt.Sprintf("The %s %s.", field, tag)
			}

			return response, err
		}

		// No validation errors, return nil
		return nil, nil
	}
}

// validateStruct performs the validation of the provided model using the validator.
// It returns a list of ErrorValidation containing validation errors.
func validateStruct(validate *validator.Validate, model interface{}) ([]ErrorValidation, error) {
	var validationErrors []ErrorValidation

	// Check for validation errors
	err := validate.Struct(model)
	if err != nil {
		// Iterate through each validation error and create an ErrorValidation
		for _, e := range err.(validator.ValidationErrors) {
			elem := ErrorValidation{
				Field : e.Field(),
				Tag   : e.Tag(),
				Param : e.Param(),
			}
			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors, err
}
