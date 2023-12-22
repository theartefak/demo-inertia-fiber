package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// ErrorResponse represents the structure for validation error responses.
type ErrorResponse struct {
	Error       bool        `json:"-"`
	FailedField string      `json:"field"`
	Tag         string      `json:"tags"`
	Value       interface{} `json:"value"`
}

var validate = validator.New()

// Validate is a middleware function for request body validation using the provided model.
// It returns a list of validation errors and an error response for invalid request bodies.
func Validate(c *fiber.Ctx, model interface{}) ([]ErrorResponse, error) {
	// Parse request body into the provided model
	if err := c.BodyParser(model); err != nil {
		return nil, c.JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validate the parsed model and return any validation errors
	if err := validateStruct(validate, model); err != nil {
		return err, nil
	}

	return nil, nil
}

// validateStruct performs the validation of the provided model using the validator.
// It returns a list of ErrorResponse containing validation errors.
func validateStruct(validate *validator.Validate, model interface{}) []ErrorResponse {
	var validationErrors []ErrorResponse

	// Check for validation errors
	if err := validate.Struct(model); err != nil {
		// Iterate through each validation error and create an ErrorResponse
		for _, e := range err.(validator.ValidationErrors) {
			elem := ErrorResponse{
				FailedField : e.Field(),
				Tag         : e.Tag(),
				Value       : e.Value(),
				Error       : true,
			}
			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}
