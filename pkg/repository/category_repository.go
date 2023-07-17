package repository

import (
	"context"

	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/pkg/domain"
)

type CategoryRepository interface {
	Insert(ctx context.Context, category domain.Category) (lastInsertId int)
	Update(ctx context.Context, category domain.Category)
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) (category domain.Category, err error)
	FindAll(ctx context.Context) (responses []domain.Category)
}
