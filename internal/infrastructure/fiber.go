package infrastructure

import (
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/exception"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
)

func NewFiber(config *viper.Viper) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      config.GetString("app.name"),
		ErrorHandler: exception.NewErrorHandler(),
		Prefork:      config.GetBool("app.prefork"),
		WriteTimeout: config.GetDuration("app.timeout"),
		ReadTimeout:  config.GetDuration("app.timeout"),
	})
	app.Use(recover.New())

	// app.Post("/api/register")

	return app
}
