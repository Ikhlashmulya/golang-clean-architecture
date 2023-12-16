package route

import (
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/delivery/http/handler"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App            *fiber.App
	UserHandler    *handler.UserHandler
	AuthMiddleware fiber.Handler
}

func RegisterRoute(app *fiber.App, userHandler *handler.UserHandler, authMiddleware fiber.Handler) *RouteConfig {
	return &RouteConfig{
		App:            app,
		UserHandler:    userHandler,
		AuthMiddleware: authMiddleware,
	}
}

func (r *RouteConfig) SetupRoute() {
	r.SetupPublicRoute()
	r.SetupProtectedRoute()
}

func (r *RouteConfig) SetupPublicRoute() {
	r.App.Post("/api/users", r.UserHandler.Register)
	r.App.Post("/api/users/_login", r.UserHandler.Login)
}

func (r *RouteConfig) SetupProtectedRoute() {
	r.App.Use(r.AuthMiddleware)
	r.App.Get("/api/users/_current", r.UserHandler.Current)
	r.App.Patch("/api/users/_current", r.UserHandler.Update)
}
