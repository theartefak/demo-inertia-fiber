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

// ErrorResponse represents the structure for a validation error response.
type ErrorResponse struct {
	Component string        `json:"component"`
	Props     ResponseProps `json:"props"`
	URL       string        `json:"url"`
	Version   string        `json:"version"`
}

// ResponseProps represents the structure for error response properties.
type ResponseProps struct {
	Errors map[string]string `json:"errors"`
}

// ErrorValidation represents a single validation error.
type ErrorValidation struct {
	Field string `json:"field"`
	Tag   string `json:"tags"`
	Param string `json:"param"`
}

var validate = validator.New()

// Validate is a middleware function for request body validation using the provided model.
// It returns a validation error response if there are errors in the request body.
func Validate(component string, c *fiber.Ctx, model interface{}) (*ErrorResponse, error) {
	// Parse request body into the provided model
	if err := c.BodyParser(model); err != nil {
		// Failed to parse request body, return error response
		return &ErrorResponse{
			Component: component,
			Props: ResponseProps{
				Errors: map[string]string{
					"body": "Failed to parse request body",
				},
			},
			URL:     strings.TrimPrefix(c.Get("Referer"), c.BaseURL()),
			Version: c.GetRespHeader("X-Inertia-Version"),
		}, err
	}

	// Validate the parsed model and return any validation errors
	validationErrors := validateStruct(validate, model)
	if len(validationErrors) > 0 {
		// Construct error response for validation errors
		response := &ErrorResponse{
			Component: component,
			Props: ResponseProps{
				Errors: make(map[string]string),
			},
			URL:     strings.TrimPrefix(c.Get("Referer"), c.BaseURL()),
			Version: c.GetRespHeader("X-Inertia-Version"),
		}

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

			response.Props.Errors[field] = fmt.Sprintf("The %s %s.", field, tag)
		}

		return response, nil
	}

	// No validation errors, return nil
	return nil, nil
}

// validateStruct performs the validation of the provided model using the validator.
// It returns a list of ErrorValidation containing validation errors.
func validateStruct(validate *validator.Validate, model interface{}) []ErrorValidation {
	var validationErrors []ErrorValidation

	// Check for validation errors
	if err := validate.Struct(model); err != nil {
		// Iterate through each validation error and create an ErrorValidation
		for _, e := range err.(validator.ValidationErrors) {
			elem := ErrorValidation{
				Field: e.Field(),
				Tag:   e.Tag(),
				Param: e.Param(),
			}
			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

func JSONResponse(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(data)
}
