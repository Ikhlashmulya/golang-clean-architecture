package infrastructure

import (
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/config"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/delivery/http/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

func NewFiberApp(categoryHandler *handler.CategoryHandler) *fiber.App {
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	app.Get("/swagger/*", swagger.HandlerDefault)

	categoryHandler.Route(app)

	return app
}
