package exception

import (
	"database/sql"
	"errors"

	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/pkg/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return ctx.
			Status(fiber.StatusNotFound).
			JSON(model.WebResponse{
				Code:    fiber.StatusNotFound,
				Status:  "NOT_FOUND",
				Message: "Data is not found",
			})
	}

	_, ok := err.(validator.ValidationErrors)
	if ok {
		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(model.WebResponse{
				Code:    fiber.StatusBadRequest,
				Status:  "BAD_REQUEST",
				Message: err.Error(),
			})
	}

	return ctx.
		Status(fiber.StatusInternalServerError).
		JSON(model.WebResponse{
			Code:    fiber.StatusInternalServerError,
			Status:  "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
		})
}
