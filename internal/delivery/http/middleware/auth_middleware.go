package middleware

import (
	"strings"

	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/model"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func NewAuth(userUsecase usecase.UserUsecase, logger *logrus.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorization := c.Get("Authorization")

		bearerToken := strings.Split(authorization, " ")
		if bearerToken[0] != "Bearer" {
			return fiber.ErrUnauthorized
		}

		auth, err := userUsecase.Verify(c.Context(), &model.VerifyUserRequest{AccessToken: bearerToken[1]})
		if err != nil {
			logger.WithError(err).Warn("user not verified")
			return err
		}

		c.Locals("auth", auth)
		
		return c.Next()
	}
}