package infrastructure

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
)

func NewFiber(config *viper.Viper) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      config.GetString("app.name"),
		ErrorHandler: NewErrorHandler(),
		Prefork:      config.GetBool("app.prefork"),
		WriteTimeout: config.GetDuration("app.timeout"),
		ReadTimeout:  config.GetDuration("app.timeout"),
	})
	app.Use(recover.New())

	// app.Post("/api/register")

	return app
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		statusCode := fiber.StatusInternalServerError

		if e, ok := err.(*fiber.Error); ok {
			statusCode = e.Code
		}

		validationErrors, ok := err.(validator.ValidationErrors)
		if ok {
			statusCode = fiber.StatusBadGateway
			errorMessages := make([]string, 0)
			for _, value := range validationErrors {
				errorMessages = append(errorMessages, fmt.Sprintf(
					"[%s]: '%v' | needs to implements '%s'",
					value.Field(),
					value.Value(),
					value.ActualTag(),
				))
			}
			
			return c.Status(statusCode).JSON(fiber.Map{
				"errors": errorMessages,
			})
		}

		return c.Status(statusCode).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
}
