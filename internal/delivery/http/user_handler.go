package http

import (
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/model"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	UserUsecase usecase.UserUsecase
	Logger      *logrus.Logger
}

func NewUserHandler(userUsecase usecase.UserUsecase, log *logrus.Logger) *UserHandler {
	return &UserHandler{
		UserUsecase: userUsecase,
		Logger:      log,
	}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	registerUserRequest := new(model.RegisterUserRequest)
	if err := c.BodyParser(registerUserRequest); err != nil {
		h.Logger.WithError(err).Error("error parsing request body")
		return err
	}

	response, err := h.UserUsecase.Register(c.Context(), registerUserRequest)
	if err != nil {
		h.Logger.WithError(err).Error("error user register")
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	loginUserReequest := new(model.LoginUserRequest)
	if err := c.BodyParser(loginUserReequest); err != nil {
		h.Logger.WithError(err).Error("error parsing request body")
		return err
	}

	response, err := h.UserUsecase.Login(c.Context(), loginUserReequest)
	if err != nil {
		h.Logger.WithError(err).Error("error user login")
		return err
	}

	return c.JSON(response)
}