package handler

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

	return c.
		Status(fiber.StatusCreated).
		JSON(&model.WebResponse[*model.UserResponse]{
			Data: response,
		})
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

	return c.
		JSON(&model.WebResponse[*model.TokenResponse]{
			Data: response,
		})
}

func (h *UserHandler) Current(c *fiber.Ctx) error {
	auth := c.Locals("auth").(*model.Auth)

	response, err := h.UserUsecase.Current(c.Context(), &model.GetUserRequest{Username: auth.Username})
	if err != nil {
		h.Logger.WithError(err).Error("error get current user")
		return err
	}

	return c.
		JSON(&model.WebResponse[*model.UserResponse]{
			Data: response,
		})
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
	auth := c.Locals("auth").(*model.Auth)

	updateUserRequest := new(model.UpdateUserRequest)
	if err := c.BodyParser(updateUserRequest); err != nil {
		h.Logger.WithError(err).Error("error parsing request body")
		return err
	}

	updateUserRequest.Username = auth.Username
	response, err := h.UserUsecase.Update(c.Context(), updateUserRequest)
	if err != nil {
		h.Logger.WithError(err).Error("error update user")
		return err
	}

	return c.
		JSON(&model.WebResponse[*model.UserResponse]{
			Data: response,
		})

}
