package usecase

import (
	"context"

	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/model"
)

type CategoryUsecase interface {
	Create(ctx context.Context, request model.CreateCategoryRequest) (response model.CategoryResponse)
	Update(ctx context.Context, request model.UpdateCategoryRequest) (response model.CategoryResponse)
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) (response model.CategoryResponse)
	FindAll(ctx context.Context) (responses []model.CategoryResponse)
}