package exception

import "github.com/gofiber/fiber/v2"

var (
	// error user domain
	ErrUserNotFound     = fiber.NewError(fiber.StatusNotFound, "User is not found")
	ErrUserAlreadyExist = fiber.NewError(fiber.StatusBadRequest, "username already exist")
	ErrPasswordNotMatch = fiber.NewError(fiber.StatusBadRequest, "password not match")

	//error
	ErrInternalServerError = fiber.ErrInternalServerError
)
