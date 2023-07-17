package test

import (
	"context"
	"errors"
	"testing"

	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/pkg/domain"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/pkg/mock"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/pkg/model"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/pkg/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateUsecase(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		categoryRepository := mock.NewMockCategoryRepository(ctrl)

		category := domain.Category{
			Name: "category1",
		}

		categoryRepository.EXPECT().Insert(context.Background(), category).Return(1)

		validate := validator.New()
		categoryUsecase := usecase.NewCategoryUsecase(validate, categoryRepository)

		categoryResponse := categoryUsecase.Create(context.Background(), model.CreateCategoryRequest{Name: "category1"})

		assert.True(t, ctrl.Satisfied())
		assert.Equal(t, 1, categoryResponse.Id)
	})
	t.Run("validation failed", func(t *testing.T) {
		assert.Panics(t, func() {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			categoryRepository := mock.NewMockCategoryRepository(ctrl)

			category := domain.Category{
				Name: "category1",
			}

			categoryRepository.EXPECT().Insert(context.Background(), category).Return(1)

			validate := validator.New()
			categoryUsecase := usecase.NewCategoryUsecase(validate, categoryRepository)

			// panic
			_ = categoryUsecase.Create(context.Background(), model.CreateCategoryRequest{})
		})
	})
}

func TestUpdateUsecase(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		categoryRepository := mock.NewMockCategoryRepository(ctrl)

		category := domain.Category{
			Id:   1,
			Name: "category2",
		}

		categoryRepository.EXPECT().FindById(context.Background(), category.Id).Return(category, nil)
		categoryRepository.EXPECT().Update(context.Background(), category)

		validate := validator.New()
		categoryUsecase := usecase.NewCategoryUsecase(validate, categoryRepository)

		categoryResponse := categoryUsecase.Update(context.Background(), model.UpdateCategoryRequest{
			Id:   category.Id,
			Name: category.Name,
		})

		assert.True(t, ctrl.Satisfied())
		assert.Equal(t, category.Id, categoryResponse.Id)
		assert.Equal(t, category.Name, categoryResponse.Name)
	})
	t.Run("not found failed", func(t *testing.T) {
		assert.Panics(t, func() {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			categoryRepository := mock.NewMockCategoryRepository(ctrl)

			category := domain.Category{
				Id:   1,
				Name: "category2",
			}

			categoryRepository.EXPECT().FindById(context.Background(), category.Id).Return(category, errors.New("category is bot found"))
			// categoryRepository.EXPECT().Update(context.Background(), category)

			validate := validator.New()
			categoryUsecase := usecase.NewCategoryUsecase(validate, categoryRepository)

			categoryResponse := categoryUsecase.Update(context.Background(), model.UpdateCategoryRequest{
				Id:   category.Id,
				Name: category.Name,
			})

			assert.True(t, ctrl.Satisfied())
			assert.Empty(t, categoryResponse)
		})
	})
	t.Run("validation failed", func(t *testing.T) {
		assert.Panics(t, func() {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			categoryRepository := mock.NewMockCategoryRepository(ctrl)

			category := domain.Category{
				Id:   1,
				Name: "category2",
			}

			categoryRepository.EXPECT().FindById(context.Background(), category.Id).Return(category, errors.New("category is bot found"))
			// categoryRepository.EXPECT().Update(context.Background(), category)

			validate := validator.New()
			categoryUsecase := usecase.NewCategoryUsecase(validate, categoryRepository)

			// panic
			_ = categoryUsecase.Update(context.Background(), model.UpdateCategoryRequest{})
		})
	})
}

func TestDeleteUsecase(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		categoryRepository := mock.NewMockCategoryRepository(ctrl)

		category := domain.Category{
			Id:   1,
			Name: "category2",
		}

		categoryRepository.EXPECT().FindById(context.Background(), category.Id).Return(category, nil)
		categoryRepository.EXPECT().Delete(context.Background(), category.Id)

		validate := validator.New()
		categoryUsecase := usecase.NewCategoryUsecase(validate, categoryRepository)

		categoryUsecase.Delete(context.Background(), category.Id)

		assert.True(t, ctrl.Satisfied())
	})

	t.Run("not found failed", func(t *testing.T) {
		assert.Panics(t, func() {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			categoryRepository := mock.NewMockCategoryRepository(ctrl)

			category := domain.Category{
				Id:   1,
				Name: "category2",
			}

			categoryRepository.EXPECT().FindById(context.Background(), category.Id).Return(category, errors.New("category is bot found"))
			// categoryRepository.EXPECT().Delete(context.Background(), category.Id)

			validate := validator.New()
			categoryUsecase := usecase.NewCategoryUsecase(validate, categoryRepository)

			categoryUsecase.Delete(context.Background(), category.Id)

			assert.True(t, ctrl.Satisfied())
		})
	})
}

func TestFindByIdUsecase(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		categoryRepository := mock.NewMockCategoryRepository(ctrl)

		category := domain.Category{
			Id:   1,
			Name: "category2",
		}

		categoryRepository.EXPECT().FindById(context.Background(), category.Id).Return(category, nil)

		validate := validator.New()
		categoryUsecase := usecase.NewCategoryUsecase(validate, categoryRepository)

		categoryResponse := categoryUsecase.FindById(context.Background(), category.Id)

		assert.True(t, ctrl.Satisfied())
		assert.Equal(t, category.Id, categoryResponse.Id)
		assert.Equal(t, category.Name, categoryResponse.Name)
	})

	t.Run("failed", func(t *testing.T) {
		assert.Panics(t, func() {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			categoryRepository := mock.NewMockCategoryRepository(ctrl)

			category := domain.Category{}

			categoryRepository.EXPECT().FindById(context.Background(), category.Id).Return(category, errors.New("category is bot found"))

			validate := validator.New()
			categoryUsecase := usecase.NewCategoryUsecase(validate, categoryRepository)

			categoryResponse := categoryUsecase.FindById(context.Background(), category.Id)

			assert.True(t, ctrl.Satisfied())
			assert.Empty(t, categoryResponse)
		})
	})
}

func TestFindAllUsecase(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		categoryRepository := mock.NewMockCategoryRepository(ctrl)

		categories := []domain.Category{
			{
				Id:   1,
				Name: "category1",
			},
			{
				Id:   2,
				Name: "category2",
			},
		}

		categoryRepository.EXPECT().FindAll(context.Background()).Return(categories)

		validate := validator.New()
		categoryUsecase := usecase.NewCategoryUsecase(validate, categoryRepository)

		categoryResponses := categoryUsecase.FindAll(context.Background())

		assert.True(t, ctrl.Satisfied())

		for i := 0; i < len(categories); i++ {
			assert.Equal(t, categories[i].Id, categoryResponses[i].Id)
			assert.Equal(t, categories[i].Name, categoryResponses[i].Name)
		}
	})

	t.Run("no content", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		categoryRepository := mock.NewMockCategoryRepository(ctrl)

		categoryRepository.EXPECT().FindAll(context.Background()).Return(nil)

		validate := validator.New()
		categoryUsecase := usecase.NewCategoryUsecase(validate, categoryRepository)

		categoryResponses := categoryUsecase.FindAll(context.Background())

		assert.True(t, ctrl.Satisfied())
		assert.Empty(t, categoryResponses)
	})
}
