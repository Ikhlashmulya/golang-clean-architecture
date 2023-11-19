package infrastructure

import (
	"time"

	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/delivery/http/handler"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/exception"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

func NewFiberApp(categoryHandler *handler.CategoryHandler) *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork:      true,
		ErrorHandler: exception.ErrorHandler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})
	app.Use(recover.New())

	app.Get("/swagger/*", swagger.HandlerDefault)

	categoryHandler.Route(app)

	return app
}
