package route

import (
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/delivery/http/handler"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoute(app *fiber.App, userHandler *handler.UserHandler, authMiddleware fiber.Handler) {
	publicRouter := app.Group("/api")
	publicRouter.Post("/users", userHandler.Register)
	publicRouter.Post("/users/_login", userHandler.Login)

	protectedRouter := app.Group("/api", authMiddleware)
	protectedRouter.Get("/users/_current", userHandler.Current)
	protectedRouter.Patch("/users/_current", userHandler.Update)
}
