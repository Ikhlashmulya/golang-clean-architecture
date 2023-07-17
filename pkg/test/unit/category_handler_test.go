package test

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/pkg/delivery/http/handler"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/pkg/mock"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/pkg/model"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	app := fiber.New()

	createCategoryRequest := model.CreateCategoryRequest{Name: "category1"}
	categoryResponse := model.CategoryResponse{
		Id:   1,
		Name: "category1",
	}

	categoryUsecase := mock.NewMockCategoryUsecase(ctrl)
	categoryUsecase.EXPECT().Create(gomock.Any(), createCategoryRequest).Return(categoryResponse)

	categoryHandler := handler.NewCategoryHandler(categoryUsecase)
	categoryHandler.Route(app)

	requestBody := strings.NewReader(`{"name": "category1"}`)

	request := httptest.NewRequest(fiber.MethodPost, "/api/categories", requestBody)
	request.Header.Set("content-type", "application/json")

	response, err := app.Test(request)
	assert.NoError(t, err)

	assert.Equal(t, 201, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)

	var responseBody model.WebResponse
	err = json.Unmarshal(body, &responseBody)
	assert.NoError(t, err)

	assert.Equal(t, 201, responseBody.Code)
	assert.Equal(t, "success create category", responseBody.Message)
	assert.Equal(t, "CREATED", responseBody.Status)

	responseBodyData := responseBody.Data.(map[string]any)
	assert.Equal(t, categoryResponse.Id, int(responseBodyData["id"].(float64)))
	assert.Equal(t, categoryResponse.Name, responseBodyData["name"])
}

func TestUpdateHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	app := fiber.New()

	category := model.UpdateCategoryRequest{
		Id:   2,
		Name: "category2",
	}

	categoryUsecase := mock.NewMockCategoryUsecase(ctrl)
	categoryUsecase.EXPECT().Update(gomock.Any(), category).Return(model.CategoryResponse(category))

	categoryHandler := handler.NewCategoryHandler(categoryUsecase)
	categoryHandler.Route(app)

	requestBody := strings.NewReader(`{"name": "category2"}`)

	request := httptest.NewRequest(fiber.MethodPut, "/api/categories/2", requestBody)
	request.Header.Set("content-type", "application/json")

	response, err := app.Test(request)
	assert.NoError(t, err)

	assert.Equal(t, 200, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)

	var responseBody model.WebResponse
	err = json.Unmarshal(body, &responseBody)
	assert.NoError(t, err)

	assert.Equal(t, 200, responseBody.Code)
	assert.Equal(t, "success update category", responseBody.Message)
	assert.Equal(t, "OK", responseBody.Status)

	responseBodyData := responseBody.Data.(map[string]any)
	assert.Equal(t, category.Id, int(responseBodyData["id"].(float64)))
	assert.Equal(t, category.Name, responseBodyData["name"])
}

func TestDeleteHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	app := fiber.New()

	categoryUsecase := mock.NewMockCategoryUsecase(ctrl)
	categoryUsecase.EXPECT().Delete(gomock.Any(), 2)

	categoryHandler := handler.NewCategoryHandler(categoryUsecase)
	categoryHandler.Route(app)

	request := httptest.NewRequest(fiber.MethodDelete, "/api/categories/2", nil)

	response, err := app.Test(request)
	assert.NoError(t, err)

	assert.Equal(t, 200, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)

	var responseBody model.WebResponse
	err = json.Unmarshal(body, &responseBody)
	assert.NoError(t, err)

	assert.Equal(t, 200, responseBody.Code)
	assert.Equal(t, "success delete category", responseBody.Message)
	assert.Equal(t, "OK", responseBody.Status)
}

func TestFindByIdHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	app := fiber.New()

	categoryResponse := model.CategoryResponse{
		Id:   2,
		Name: "category2",
	}

	categoryUsecase := mock.NewMockCategoryUsecase(ctrl)
	categoryUsecase.EXPECT().FindById(gomock.Any(), categoryResponse.Id).Return(categoryResponse)

	categoryHandler := handler.NewCategoryHandler(categoryUsecase)
	categoryHandler.Route(app)

	request := httptest.NewRequest(fiber.MethodGet, "/api/categories/2", nil)

	response, err := app.Test(request)
	assert.NoError(t, err)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)

	var responseBody model.WebResponse
	err = json.Unmarshal(body, &responseBody)
	assert.NoError(t, err)

	assert.Equal(t, 200, responseBody.Code)
	assert.Equal(t, "success get category", responseBody.Message)
	assert.Equal(t, "OK", responseBody.Status)

	responseBodyData := responseBody.Data.(map[string]any)
	assert.Equal(t, categoryResponse.Id, int(responseBodyData["id"].(float64)))
	assert.Equal(t, categoryResponse.Name, responseBodyData["name"])
}

func TestFindAllHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	app := fiber.New()

	categoryResponses := []model.CategoryResponse{
		{
			Id:   1,
			Name: "category1",
		},
		{
			Id:   2,
			Name: "category2",
		},
	}

	categoryUsecase := mock.NewMockCategoryUsecase(ctrl)
	categoryUsecase.EXPECT().FindAll(gomock.Any()).Return(categoryResponses)

	categoryHandler := handler.NewCategoryHandler(categoryUsecase)
	categoryHandler.Route(app)

	request := httptest.NewRequest(fiber.MethodGet, "/api/categories", nil)

	response, err := app.Test(request)
	assert.NoError(t, err)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)

	var responseBody model.WebResponse
	err = json.Unmarshal(body, &responseBody)
	assert.NoError(t, err)

	assert.Equal(t, 200, responseBody.Code)
	assert.Equal(t, "success get all category", responseBody.Message)
	assert.Equal(t, "OK", responseBody.Status)

	responseBodyData := responseBody.Data.([]any)

	for i := 0; i < len(responseBodyData); i++ {
		data := responseBodyData[i].(map[string]any)
		assert.Equal(t, categoryResponses[i].Id, int(data["id"].(float64)))
		assert.Equal(t, categoryResponses[i].Name, data["name"])
	}
}
