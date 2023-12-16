package integration

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/model"
)

func (s *e2eTestSuite) TestUserRegisterSuccess() {
	requestBody := &model.RegisterUserRequest{
		Name:     "John Doe",
		Username: "johndoe",
		Password: "johndoe123",
	}

	bodyJSON, err := json.Marshal(requestBody)
	s.Assert().NoError(err)

	request := httptest.NewRequest(http.MethodPost, "/api/users", strings.NewReader(string(bodyJSON)))
	request.Header.Add("content-type", "application/json")

	response, err := s.App.Test(request)
	s.Assert().NoError(err)
	s.Assert().Equal(http.StatusCreated, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	s.Assert().NoError(err)

	responseBody := new(model.WebResponse[*model.UserResponse])
	err = json.Unmarshal(bytes, responseBody)
	s.Assert().NoError(err)

	s.Assert().Equal(requestBody.Name, responseBody.Data.Name)
}

func (s *e2eTestSuite) TestUserRegisterFailedValidation() {
	requestBody := &model.RegisterUserRequest{
		Name:     "",
		Username: "",
		Password: "",
	}

	bodyJSON, err := json.Marshal(requestBody)
	s.Assert().NoError(err)

	request := httptest.NewRequest(http.MethodPost, "/api/users", strings.NewReader(string(bodyJSON)))
	request.Header.Add("content-type", "application/json")

	response, err := s.App.Test(request)
	s.Assert().NoError(err)
	s.Assert().Equal(http.StatusBadRequest, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	s.Assert().NoError(err)

	responseBody := make(map[string]any)
	err = json.Unmarshal(bytes, &responseBody)
	s.Assert().NoError(err)

	s.Assert().NotEmpty(responseBody["errors"])
}

func (s *e2eTestSuite) TestUserRegisterFailedUserAlreadyExists() {
	s.TestUserRegisterSuccess()

	requestBody := &model.RegisterUserRequest{
		Name:     "John Doe",
		Username: "johndoe",
		Password: "johndoe123",
	}

	bodyJSON, err := json.Marshal(requestBody)
	s.Assert().NoError(err)

	request := httptest.NewRequest(http.MethodPost, "/api/users", strings.NewReader(string(bodyJSON)))
	request.Header.Add("content-type", "application/json")

	response, err := s.App.Test(request)
	s.Assert().NoError(err)
	s.Assert().Equal(http.StatusBadRequest, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	s.Assert().NoError(err)

	responseBody := make(map[string]any)
	err = json.Unmarshal(bytes, &responseBody)
	s.Assert().NoError(err)

	s.Assert().NotEmpty(responseBody["errors"])
}

func (s *e2eTestSuite) TestUserLoginSuccess() {
	s.TestUserRegisterSuccess()

	requestBody := &model.LoginUserRequest{
		Username: "johndoe",
		Password: "johndoe123",
	}

	bodyJSON, err := json.Marshal(requestBody)
	s.Assert().NoError(err)

	request := httptest.NewRequest(http.MethodPost, "/api/users/_login", strings.NewReader(string(bodyJSON)))
	request.Header.Add("content-type", "application/json")

	response, err := s.App.Test(request)
	s.Assert().NoError(err)
	s.Assert().Equal(http.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	s.Assert().NoError(err)

	responseBody := new(model.WebResponse[*model.TokenResponse])
	err = json.Unmarshal(bytes, responseBody)
	s.Assert().NoError(err)

	s.Assert().NotEmpty(responseBody.Data.AccessToken)
	s.Assert().Equal("Bearer", responseBody.Data.TokenType)
}

func (s *e2eTestSuite) TestUserLoginFailedUserNotFound() {
	s.TestUserRegisterSuccess()

	requestBody := &model.LoginUserRequest{
		Username: "wrongjohndoe",
		Password: "johndoe123",
	}

	bodyJSON, err := json.Marshal(requestBody)
	s.Assert().NoError(err)

	request := httptest.NewRequest(http.MethodPost, "/api/users/_login", strings.NewReader(string(bodyJSON)))
	request.Header.Add("content-type", "application/json")

	response, err := s.App.Test(request)
	s.Assert().NoError(err)
	s.Assert().Equal(http.StatusNotFound, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	s.Assert().NoError(err)

	responseBody := make(map[string]any)
	err = json.Unmarshal(bytes, &responseBody)
	s.Assert().NoError(err)

	s.Assert().NotEmpty(responseBody["errors"])
}

func (s *e2eTestSuite) TestUserLoginFailedPasswordNotMatch() {
	s.TestUserRegisterSuccess()

	requestBody := &model.LoginUserRequest{
		Username: "johndoe",
		Password: "wrongpassword",
	}

	bodyJSON, err := json.Marshal(requestBody)
	s.Assert().NoError(err)

	request := httptest.NewRequest(http.MethodPost, "/api/users/_login", strings.NewReader(string(bodyJSON)))
	request.Header.Add("content-type", "application/json")

	response, err := s.App.Test(request)
	s.Assert().NoError(err)
	s.Assert().Equal(http.StatusBadRequest, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	s.Assert().NoError(err)

	responseBody := make(map[string]any)
	err = json.Unmarshal(bytes, &responseBody)
	s.Assert().NoError(err)

	s.Assert().NotEmpty(responseBody["errors"])
}

func (s *e2eTestSuite) TestUserLoginFailedValidation() {
	s.TestUserRegisterSuccess()

	requestBody := &model.LoginUserRequest{
		Username: "",
		Password: "",
	}

	bodyJSON, err := json.Marshal(requestBody)
	s.Assert().NoError(err)

	request := httptest.NewRequest(http.MethodPost, "/api/users/_login", strings.NewReader(string(bodyJSON)))
	request.Header.Add("content-type", "application/json")

	response, err := s.App.Test(request)
	s.Assert().NoError(err)
	s.Assert().Equal(http.StatusBadRequest, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	s.Assert().NoError(err)

	responseBody := make(map[string]any)
	err = json.Unmarshal(bytes, &responseBody)
	s.Assert().NoError(err)

	s.Assert().NotEmpty(responseBody["errors"])
}

func (s *e2eTestSuite) TestUserCurrentSuccess() {
	token := s.GetTokenUser()

	request := httptest.NewRequest(http.MethodGet, "/api/users/_current", nil)
	request.Header.Add("content-type", "application/json")
	request.Header.Add("Authorization", "Bearer "+token)

	response, err := s.App.Test(request)
	s.Assert().NoError(err)
	s.Assert().Equal(http.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	s.Assert().NoError(err)

	responseBody := new(model.WebResponse[*model.UserResponse])
	err = json.Unmarshal(bytes, responseBody)
	s.Assert().NoError(err)

	s.Assert().Equal("John Doe", responseBody.Data.Name)
	s.Assert().Equal("johndoe", responseBody.Data.Username)
	s.Assert().NotEmpty(responseBody.Data.CreatedAt)
	s.Assert().NotEmpty(responseBody.Data.UpdatedAt)
}

func (s *e2eTestSuite) TestUserCurrentFailedUnauthorized() {
	request := httptest.NewRequest(http.MethodGet, "/api/users/_current", nil)
	request.Header.Add("content-type", "application/json")

	response, err := s.App.Test(request)
	s.Assert().NoError(err)
	s.Assert().Equal(http.StatusUnauthorized, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	s.Assert().NoError(err)

	responseBody := make(map[string]any)
	err = json.Unmarshal(bytes, &responseBody)
	s.Assert().NoError(err)

	s.Assert().NotEmpty(responseBody["errors"])
}

func (s *e2eTestSuite) TestUserUpdateSuccess() {
	token := s.GetTokenUser()

	requestBody := &model.UpdateUserRequest{
		Name:     "John Doe Update",
		Password: "johndoeupdate",
	}

	bodyJSON, err := json.Marshal(requestBody)
	s.Assert().NoError(err)

	request := httptest.NewRequest(http.MethodPatch, "/api/users/_current", strings.NewReader(string(bodyJSON)))
	request.Header.Add("content-type", "application/json")
	request.Header.Add("Authorization", "Bearer "+token)

	response, err := s.App.Test(request)
	s.Assert().NoError(err)
	s.Assert().Equal(http.StatusOK, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	s.Assert().NoError(err)

	responseBody := new(model.WebResponse[*model.UserResponse])
	err = json.Unmarshal(bytes, responseBody)
	s.Assert().NoError(err)

	s.Assert().Equal(requestBody.Name, responseBody.Data.Name)
}

func (s *e2eTestSuite) TestUserUpdateFailedUnauthorized() {
	requestBody := &model.UpdateUserRequest{
		Name:     "John Doe Update",
		Password: "johndoeupdate",
	}

	bodyJSON, err := json.Marshal(requestBody)
	s.Assert().NoError(err)

	request := httptest.NewRequest(http.MethodPatch, "/api/users/_current", strings.NewReader(string(bodyJSON)))
	request.Header.Add("content-type", "application/json")

	response, err := s.App.Test(request)
	s.Assert().NoError(err)
	s.Assert().Equal(http.StatusUnauthorized, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	s.Assert().NoError(err)

	responseBody := make(map[string]any)
	err = json.Unmarshal(bytes, &responseBody)
	s.Assert().NoError(err)

	s.Assert().NotEmpty(responseBody["errors"])
}

func (s *e2eTestSuite) GetTokenUser() string {
	s.TestUserRegisterSuccess()
	tokenResponse, err := s.UserUsecase.Login(context.Background(), &model.LoginUserRequest{Username: "johndoe", Password: "johndoe123"})
	s.Assert().NoError(err)

	return tokenResponse.AccessToken
}
