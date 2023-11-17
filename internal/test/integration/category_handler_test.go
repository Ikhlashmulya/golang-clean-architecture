package test

import (
	"context"
	"encoding/json"
	"io"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/config"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/delivery/http/handler"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/domain"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/infrastructure"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/model"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/repository"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/stretchr/testify/assert"
)

var configuration = config.NewConfig(".")
var db = infrastructure.NewDB(configuration)
var validate = validator.New()
var categoryRepository = repository.NewCategoryRepository(db)
var categoryUsecase = usecase.NewCategoryUsecase(validate, categoryRepository)
var categoryHandler = handler.NewCategoryHandler(categoryUsecase)
var app = setupApp()

func setupApp() *fiber.App {
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	categoryHandler.Route(app)

	return app
}

func TestCreateCategory(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		requestBody := strings.NewReader(`{"name": "category1"}`)
		request := httptest.NewRequest(fiber.MethodPost, "/api/categories", requestBody)
		request.Header.Add("content-type", "application/json")

		response, err := app.Test(request)
		assert.NoError(t, err)

		assert.Equal(t, 201, response.StatusCode)

		body, err := io.ReadAll(response.Body)
		assert.NoError(t, err)

		var responseBody model.WebResponse
		err = json.Unmarshal(body, &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, fiber.StatusCreated, responseBody.Code)
		assert.Equal(t, "CREATED", responseBody.Status)
		assert.Equal(t, "success create category", responseBody.Message)

		responseBodyData := responseBody.Data.(map[string]any)
		assert.Equal(t, "category1", responseBodyData["name"])

		categoryRepository.Delete(context.Background(), int(responseBodyData["id"].(float64)))
	})
	t.Run("BAD_REQUEST", func(t *testing.T) {
		requestBody := strings.NewReader(`{"name": ""}`)
		request := httptest.NewRequest(fiber.MethodPost, "/api/categories", requestBody)
		request.Header.Add("content-type", "application/json")

		response, err := app.Test(request)
		assert.NoError(t, err)

		assert.Equal(t, 400, response.StatusCode)

		body, err := io.ReadAll(response.Body)
		assert.NoError(t, err)

		var responseBody model.WebResponse
		err = json.Unmarshal(body, &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, fiber.StatusBadRequest, responseBody.Code)
		assert.Equal(t, "BAD_REQUEST", responseBody.Status)
	})
}

func TestUpdateCategory(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		id := categoryRepository.Insert(context.Background(), domain.Category{Name: "category1"})

		requestBody := strings.NewReader(`{"name": "category1(edit)"}`)
		request := httptest.NewRequest(fiber.MethodPut, "/api/categories/"+strconv.Itoa(id), requestBody)
		request.Header.Add("content-type", "application/json")

		response, err := app.Test(request)
		assert.NoError(t, err)

		assert.Equal(t, 200, response.StatusCode)

		body, err := io.ReadAll(response.Body)
		assert.NoError(t, err)

		var responseBody model.WebResponse
		err = json.Unmarshal(body, &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, fiber.StatusOK, responseBody.Code)
		assert.Equal(t, "OK", responseBody.Status)
		assert.Equal(t, "success update category", responseBody.Message)

		responseBodyData := responseBody.Data.(map[string]any)
		assert.Equal(t, id, int(responseBodyData["id"].(float64)))
		assert.Equal(t, "category1(edit)", responseBodyData["name"])

		categoryRepository.Delete(context.Background(), id)
	})
	t.Run("BAD_REQUEST", func(t *testing.T) {
		requestBody := strings.NewReader(`{"name": ""}`)
		request := httptest.NewRequest(fiber.MethodPut, "/api/categories/1", requestBody)
		request.Header.Add("content-type", "application/json")

		response, err := app.Test(request)
		assert.NoError(t, err)

		assert.Equal(t, 400, response.StatusCode)

		body, err := io.ReadAll(response.Body)
		assert.NoError(t, err)

		var responseBody model.WebResponse
		err = json.Unmarshal(body, &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, fiber.StatusBadRequest, responseBody.Code)
		assert.Equal(t, "BAD_REQUEST", responseBody.Status)
	})
}

func TestDeleteCategory(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		id := categoryRepository.Insert(context.Background(), domain.Category{Name: "category1"})

		request := httptest.NewRequest(fiber.MethodDelete, "/api/categories/"+strconv.Itoa(id), nil)

		response, err := app.Test(request)
		assert.NoError(t, err)

		assert.Equal(t, 200, response.StatusCode)

		body, err := io.ReadAll(response.Body)
		assert.NoError(t, err)

		var responseBody model.WebResponse
		err = json.Unmarshal(body, &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, fiber.StatusOK, responseBody.Code)
		assert.Equal(t, "OK", responseBody.Status)
		assert.Equal(t, "success delete category", responseBody.Message)
	})

	t.Run("NOT_FOUND", func(t *testing.T) {
		request := httptest.NewRequest(fiber.MethodDelete, "/api/categories/25", nil)

		response, err := app.Test(request)
		assert.NoError(t, err)

		assert.Equal(t, 404, response.StatusCode)

		body, err := io.ReadAll(response.Body)
		assert.NoError(t, err)

		var responseBody model.WebResponse
		err = json.Unmarshal(body, &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, fiber.StatusNotFound, responseBody.Code)
		assert.Equal(t, "NOT_FOUND", responseBody.Status)
	})
}

func TestFindByIdCategory(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		id := categoryRepository.Insert(context.Background(), domain.Category{Name: "category1"})

		request := httptest.NewRequest(fiber.MethodGet, "/api/categories/"+strconv.Itoa(id), nil)

		response, err := app.Test(request)
		assert.NoError(t, err)

		assert.Equal(t, 200, response.StatusCode)

		body, err := io.ReadAll(response.Body)
		assert.NoError(t, err)

		var responseBody model.WebResponse
		err = json.Unmarshal(body, &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, fiber.StatusOK, responseBody.Code)
		assert.Equal(t, "OK", responseBody.Status)
		assert.Equal(t, "success get category", responseBody.Message)

		responseBodyData := responseBody.Data.(map[string]any)
		assert.Equal(t, id, int(responseBodyData["id"].(float64)))
		assert.Equal(t, "category1", responseBodyData["name"])

		categoryRepository.Delete(context.Background(), id)
	})

	t.Run("NOT_FOUND", func(t *testing.T) {
		request := httptest.NewRequest(fiber.MethodGet, "/api/categories/25", nil)

		response, err := app.Test(request)
		assert.NoError(t, err)

		assert.Equal(t, 404, response.StatusCode)

		body, err := io.ReadAll(response.Body)
		assert.NoError(t, err)

		var responseBody model.WebResponse
		err = json.Unmarshal(body, &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, fiber.StatusNotFound, responseBody.Code)
		assert.Equal(t, "NOT_FOUND", responseBody.Status)
	})
}

func TestFindAllCategory(t *testing.T) {
	categories := []domain.Category{
		{
			Name: "Category1",
		},
		{
			Name: "Category2",
		},
	}

	for _, category := range categories {
		categoryRepository.Insert(context.Background(), category)
	}

	request := httptest.NewRequest(fiber.MethodGet, "/api/categories", nil)

	response, err := app.Test(request)
	assert.NoError(t, err)

	assert.Equal(t, 200, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)

	var responseBody model.WebResponse
	err = json.Unmarshal(body, &responseBody)
	assert.NoError(t, err)

	assert.Equal(t, fiber.StatusOK, responseBody.Code)
	assert.Equal(t, "OK", responseBody.Status)
	assert.Equal(t, "success get all category", responseBody.Message)

	data := responseBody.Data.([]any)

	for i := 0; i < len(data); i++ {
		assert.Equal(t, categories[i].Name, data[i].(map[string]any)["name"])
	}
}

func TestMain(m *testing.M) {
	_, _ = db.Exec(`TRUNCATE category`)
	m.Run()
	_, _ = db.Exec(`TRUNCATE category`)
}
