package exception

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

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

	validationException, ok := err.(validator.ValidationErrors)
	if ok {
		errorMessages := make([]string, 0)
		for _, value := range validationException {
			errorMessages = append(errorMessages, fmt.Sprintf(
				"[%s]: '%v' | needs to implements '%s'",
				value.Field(),
				value.Value(),
				value.ActualTag(),
			))
		}
		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(model.WebResponse{
				Code:    fiber.StatusBadRequest,
				Status:  "BAD_REQUEST",
				Message: strings.Join(errorMessages, " and "),
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
