package usecase

import (
	"context"

	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/pkg/domain"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/pkg/exception"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/pkg/mapper"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/pkg/model"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/pkg/repository"
	"github.com/go-playground/validator/v10"
)

type CategoryUsecaseImpl struct {
	CategoryRepository repository.CategoryRepository
	Validate           *validator.Validate
}

func NewCategoryUsecase(validate *validator.Validate, categoryRepository repository.CategoryRepository) CategoryUsecase {
	return &CategoryUsecaseImpl{
		Validate:           validate,
		CategoryRepository: categoryRepository,
	}
}

func (usecase *CategoryUsecaseImpl) Create(ctx context.Context, request model.CreateCategoryRequest) (response model.CategoryResponse) {
	err := usecase.Validate.Struct(request)
	exception.PanicIfError(err)

	category := domain.Category{
		Name: request.Name,
	}

	lastInsertId := usecase.CategoryRepository.Insert(ctx, category)

	category.Id = lastInsertId

	response = mapper.ToCategoryResponse(category)
	return
}

func (usecase *CategoryUsecaseImpl) Update(ctx context.Context, request model.UpdateCategoryRequest) (response model.CategoryResponse) {
	err := usecase.Validate.Struct(request)
	exception.PanicIfError(err)

	category, err := usecase.CategoryRepository.FindById(ctx, request.Id)
	exception.PanicIfError(err)

	category.Name = request.Name

	usecase.CategoryRepository.Update(ctx, category)

	response = mapper.ToCategoryResponse(category)
	return
}

func (usecase *CategoryUsecaseImpl) Delete(ctx context.Context, categoryId int) {
	category, err := usecase.CategoryRepository.FindById(ctx, categoryId)
	exception.PanicIfError(err)

	usecase.CategoryRepository.Delete(ctx, category.Id)
}

func (usecase *CategoryUsecaseImpl) FindById(ctx context.Context, categoryId int) (response model.CategoryResponse) {
	category, err := usecase.CategoryRepository.FindById(ctx, categoryId)
	exception.PanicIfError(err)

	response = mapper.ToCategoryResponse(category)
	return
}

func (usecase *CategoryUsecaseImpl) FindAll(ctx context.Context) (responses []model.CategoryResponse) {
	categories := usecase.CategoryRepository.FindAll(ctx)

	for _, category := range categories {
		responses = append(responses, mapper.ToCategoryResponse(category))
	}

	return
}
