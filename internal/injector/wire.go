//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/config"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/infrastructure"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/delivery/http/handler"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/repository"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

func InitializedCategoryHandler(config *config.Config) *handler.CategoryHandler {
	wire.Build(
		infrastructure.NewDB,
		repository.NewCategoryRepository,
		validator.New,
		usecase.NewCategoryUsecase,
		handler.NewCategoryHandler,
	)

	return nil
}
