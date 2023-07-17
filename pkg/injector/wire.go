//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/config"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/pkg/database"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/pkg/delivery/http/handler"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/pkg/repository"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/pkg/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

func InitializedCategoryHandler(config *config.Config) *handler.CategoryHandler {
	wire.Build(
		database.NewDB,
		repository.NewCategoryRepository,
		validator.New,
		usecase.NewCategoryUsecase,
		handler.NewCategoryHandler,
	)

	return nil
}
