package config

import (
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/exception"
	"github.com/gofiber/fiber/v2"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}
