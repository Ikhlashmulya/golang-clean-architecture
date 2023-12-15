package route

import (
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/delivery/http/handler"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	Router         fiber.Router
	UserHandler    *handler.UserHandler
	AuthMiddleware fiber.Handler
}

func RegisterRoute(router fiber.Router, userHandler *handler.UserHandler, authMiddleware fiber.Handler) *RouteConfig {
	return &RouteConfig{
		Router:         router,
		UserHandler:    userHandler,
		AuthMiddleware: authMiddleware,
	}
}

func (r *RouteConfig) SetupRoute() {
	r.Router.Post("/users", r.UserHandler.Register)
	r.Router.Post("/users/_login", r.UserHandler.Login)
	r.Router.Use(r.AuthMiddleware)
	r.Router.Get("/users/_current", r.UserHandler.Current)
	r.Router.Patch("/users/_current", r.UserHandler.Update)
}
