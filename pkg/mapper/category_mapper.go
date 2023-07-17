package mapper

import (
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/pkg/domain"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/pkg/model"
)

func ToCategoryResponse(category domain.Category) model.CategoryResponse {
	return model.CategoryResponse{
		Name: category.Name,
		Id:   category.Id,
	}
}
