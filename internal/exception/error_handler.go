package exception

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// define error here
var (
	// error user domain
	ErrUserNotFound         = fiber.NewError(fiber.StatusNotFound, "User is not found")
	ErrUserAlreadyExist     = fiber.NewError(fiber.StatusBadRequest, "username already exist")
	ErrUserPasswordNotMatch = fiber.NewError(fiber.StatusBadRequest, "password not match")
	ErrUserUnauthorized     = fiber.NewError(fiber.StatusUnauthorized, "User unauthorized")

	//error
	ErrInternalServerError = fiber.ErrInternalServerError
)

func NewErrorHandler() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		statusCode := fiber.StatusInternalServerError
		var message any

		if e, ok := err.(*fiber.Error); ok {
			statusCode = e.Code
			message = e.Error()
		}

		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			errorMessages := make([]string, 0)
			for _, value := range validationErrors {
				errorMessages = append(errorMessages, fmt.Sprintf(
					"[%s]: '%v' | needs to implements '%s'",
					value.Field(),
					value.Value(),
					value.ActualTag(),
				))
			}

			statusCode = fiber.StatusBadRequest
			message = errorMessages
		}

		return c.Status(statusCode).JSON(fiber.Map{
			"errors": message,
		})
	}
}
