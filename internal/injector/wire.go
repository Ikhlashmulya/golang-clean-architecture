//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/delivery/http/handler"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/infrastructure"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/repository"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func InitializedApp(config *viper.Viper) *fiber.App {
	wire.Build(
		infrastructure.NewDB,
		repository.NewCategoryRepository,
		validator.New,
		usecase.NewCategoryUsecase,
		handler.NewCategoryHandler,
		infrastructure.NewFiberApp,
	)

	return nil
}
