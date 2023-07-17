package handler

import (
	"strconv"

	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/pkg/exception"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/pkg/model"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/pkg/usecase"
	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	CategoryUsecase usecase.CategoryUsecase
}

func NewCategoryHandler(categoryUsecase usecase.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{
		CategoryUsecase: categoryUsecase,
	}
}

func (handler *CategoryHandler) Route(app *fiber.App) {
	categories := app.Group("/api/categories")
	categories.Post("/", handler.Create)
	categories.Put("/:id", handler.Update)
	categories.Delete("/:id", handler.Delete)
	categories.Get("/:id", handler.FindById)
	categories.Get("/", handler.FindAll)
}

// Create Category godoc
//	@Summary		create category
//	@Description	create new category
//	@Tags			category
//	@Accept			json
//	@Produce		json
//	@Param			request	body		model.CreateCategoryRequest	true	"request body category"
//	@Success		201		{object}	model.WebResponse
//	@Failure		400		{object}	model.WebResponse
//	@Failure		500		{object}	model.WebResponse
//	@Router			/categories [post]
func (handler *CategoryHandler) Create(ctx *fiber.Ctx) error {
	var createCategoryRequest model.CreateCategoryRequest
	err := ctx.BodyParser(&createCategoryRequest)
	exception.PanicIfError(err)

	categoryResponse := handler.CategoryUsecase.Create(ctx.Context(), createCategoryRequest)

	return ctx.
		Status(fiber.StatusCreated).
		JSON(model.WebResponse{
			Code:    fiber.StatusCreated,
			Status:  "CREATED",
			Message: "success create category",
			Data:    categoryResponse,
		})
}

// Update Category godoc
//	@Summary		update category
//	@Description	update category
//	@Tags			category
//	@Accept			json
//	@Produce		json
//	@Param			request	body		model.UpdateCategoryRequest	true	"request body category"
//	@Param			id		path		int							true	"Category ID"
//	@Success		200		{object}	model.WebResponse
//	@Failure		404		{object}	model.WebResponse
//	@Failure		400		{object}	model.WebResponse
//	@Failure		500		{object}	model.WebResponse
//	@Router			/categories/{id} [put]
func (hanler *CategoryHandler) Update(ctx *fiber.Ctx) error {
	var updateCategoryRequest model.UpdateCategoryRequest
	err := ctx.BodyParser(&updateCategoryRequest)
	exception.PanicIfError(err)

	categoryId, err := strconv.Atoi(ctx.Params("id"))
	exception.PanicIfError(err)

	updateCategoryRequest.Id = categoryId

	categoryResponse := hanler.CategoryUsecase.Update(ctx.Context(), updateCategoryRequest)

	return ctx.
		Status(fiber.StatusOK).
		JSON(model.WebResponse{
			Code:    fiber.StatusOK,
			Status:  "OK",
			Message: "success update category",
			Data:    categoryResponse,
		})
}

// Delete Category godoc
//	@Summary		delete category
//	@Description	delete category
//	@Tags			category
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Category ID"
//	@Success		200	{object}	model.WebResponse
//	@Failure		404	{object}	model.WebResponse
//	@Failure		500	{object}	model.WebResponse
//	@Router			/categories/{id} [delete]
func (handler *CategoryHandler) Delete(ctx *fiber.Ctx) error {
	categoryId, err := strconv.Atoi(ctx.Params("id"))
	exception.PanicIfError(err)

	handler.CategoryUsecase.Delete(ctx.Context(), categoryId)

	return ctx.
		Status(fiber.StatusOK).
		JSON(model.WebResponse{
			Code:    fiber.StatusOK,
			Status:  "OK",
			Message: "success delete category",
		})
}

// Get Category godoc
//	@Summary		Get category
//	@Description	Get category
//	@Tags			category
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Category ID"
//	@Success		200	{object}	model.WebResponse
//	@Failure		404	{object}	model.WebResponse
//	@Failure		500	{object}	model.WebResponse
//	@Router			/categories/{id} [get]
func (handler *CategoryHandler) FindById(ctx *fiber.Ctx) error {
	categoryId, err := strconv.Atoi(ctx.Params("id"))
	exception.PanicIfError(err)

	categoryResponse := handler.CategoryUsecase.FindById(ctx.Context(), categoryId)

	return ctx.
		Status(fiber.StatusOK).
		JSON(model.WebResponse{
			Code:    fiber.StatusOK,
			Status:  "OK",
			Message: "success get category",
			Data:    categoryResponse,
		})
}

// Get Category godoc
//	@Summary		Get category
//	@Description	Get category
//	@Tags			category
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.WebResponse
//	@Failure		500	{object}	model.WebResponse
//	@Router			/categories [get]
func (handler *CategoryHandler) FindAll(ctx *fiber.Ctx) error {
	categoryResponses := handler.CategoryUsecase.FindAll(ctx.Context())

	return ctx.
		Status(fiber.StatusOK).
		JSON(model.WebResponse{
			Code:    fiber.StatusOK,
			Status:  "OK",
			Message: "success get all category",
			Data:    categoryResponses,
		})
}
