package unit

import (
	"context"
	"testing"

	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/config"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/domain"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/exception"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/model"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/usecase"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/test/unit/mocks"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ctx = context.Background()
)

func TestRegisterUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	userRepository := mocks.NewMockUserRepository(ctrl)
	userUsecase := usecase.NewUserUsecase(userRepository, logrus.New(), validator.New(), config.New())

	t.Run("success", func(t *testing.T) {
		userRepository.EXPECT().CountByUsername(ctx, "johndoe").Return(int64(0), nil)
		userRepository.EXPECT().Create(ctx, gomock.Any()).Return(nil)

		request := &model.RegisterUserRequest{
			Name:     "John Doe",
			Username: "johndoe",
			Password: "password",
		}

		response, err := userUsecase.Register(ctx, request)
		assert.NoError(t, err)
		assert.Equal(t, request.Name, response.Name)
		assert.Equal(t, request.Username, response.Username)
	})

	t.Run("failed validation", func(t *testing.T) {
		request := &model.RegisterUserRequest{
			Name:     "",
			Username: "",
			Password: "",
		}

		_, err := userUsecase.Register(ctx, request)
		assert.Error(t, err)
	})

	t.Run("failed username already exist", func(t *testing.T) {
		userRepository.EXPECT().CountByUsername(ctx, "johndoe").Return(int64(1), nil)

		request := &model.RegisterUserRequest{
			Name:     "John Doe",
			Username: "johndoe",
			Password: "password",
		}

		_, err := userUsecase.Register(ctx, request)
		assert.Error(t, err)
	})
}

func TestLoginUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	userRepository := mocks.NewMockUserRepository(ctrl)
	userUsecase := usecase.NewUserUsecase(userRepository, logrus.New(), validator.New(), config.New())

	user := createUser(t)

	t.Run("success", func(t *testing.T) {
		userRepository.EXPECT().FindByUsername(ctx, "johndoe").Return(user, nil)

		response, err := userUsecase.Login(ctx, &model.LoginUserRequest{
			Username: "johndoe",
			Password: "password",
		})

		assert.NoError(t, err)
		assert.Equal(t, "Bearer", response.TokenType)
		assert.NotEmpty(t, response.TokenType)
	})

	t.Run("failed user not found", func(t *testing.T) {
		userRepository.EXPECT().FindByUsername(ctx, "johndoe").Return(nil, gorm.ErrRecordNotFound)

		_, err := userUsecase.Login(ctx, &model.LoginUserRequest{
			Username: "johndoe",
			Password: "password",
		})

		assert.Error(t, err)
		assert.ErrorIs(t, exception.ErrUserNotFound, err)
	})

	t.Run("failed password not match", func(t *testing.T) {
		userRepository.EXPECT().FindByUsername(ctx, "johndoe").Return(user, nil)

		_, err := userUsecase.Login(ctx, &model.LoginUserRequest{
			Username: "johndoe",
			Password: "wrongPassword",
		})

		assert.Error(t, err)
		assert.ErrorIs(t, exception.ErrUserPasswordNotMatch, err)
	})

	t.Run("failed validation", func(t *testing.T) {
		_, err := userUsecase.Login(ctx, &model.LoginUserRequest{
			Username: "",
			Password: "",
		})

		assert.Error(t, err)
	})
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	userRepository := mocks.NewMockUserRepository(ctrl)
	userUsecase := usecase.NewUserUsecase(userRepository, logrus.New(), validator.New(), config.New())

	user := createUser(t)

	t.Run("success", func(t *testing.T) {
		userRepository.EXPECT().FindByUsername(ctx, "johndoe").Return(user, nil)
		userRepository.EXPECT().Update(ctx, gomock.Any()).Return(nil)

		request := &model.UpdateUserRequest{
			Username: "johndoe",
			Name:     "John Doe edited",
		}

		response, err := userUsecase.Update(ctx, request)
		assert.NoError(t, err)
		assert.Equal(t, request.Name, response.Name)
	})

	t.Run("failed user not found", func(t *testing.T) {
		userRepository.EXPECT().FindByUsername(ctx, "johndoe").Return(nil, gorm.ErrRecordNotFound)

		request := &model.UpdateUserRequest{
			Username: "johndoe",
			Name:     "John Doe edited",
		}

		_, err := userUsecase.Update(ctx, request)
		assert.Error(t, err)
		assert.ErrorIs(t, exception.ErrUserNotFound, err)

	})
}

func TestCurrentUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	userRepository := mocks.NewMockUserRepository(ctrl)
	userUsecase := usecase.NewUserUsecase(userRepository, logrus.New(), validator.New(), config.New())

	user := createUser(t)

	t.Run("success", func(t *testing.T) {
		userRepository.EXPECT().FindByUsername(ctx, "johndoe").Return(user, nil)

		response, err := userUsecase.Current(ctx, &model.GetUserRequest{Username: "johndoe"})
		assert.NoError(t, err)
		assert.Equal(t, user.ID, response.ID)
		assert.Equal(t, user.Name, response.Name)
		assert.Equal(t, user.Username, response.Username)
		assert.Equal(t, user.CreatedAt, response.CreatedAt)
		assert.Equal(t, user.UpdatedAt, response.UpdatedAt)
	})

	t.Run("failed not found", func(t *testing.T) {
		userRepository.EXPECT().FindByUsername(ctx, "johndoe").Return(nil, gorm.ErrRecordNotFound)

		_, err := userUsecase.Current(ctx, &model.GetUserRequest{Username: "johndoe"})
		assert.Error(t, err)
		assert.ErrorIs(t, exception.ErrUserNotFound, err)
	})
}

func TestVerifyUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	userRepository := mocks.NewMockUserRepository(ctrl)
	userUsecase := usecase.NewUserUsecase(userRepository, logrus.New(), validator.New(), config.New())

	user := createUser(t)

	t.Run("success", func(t *testing.T) {
		userRepository.EXPECT().FindByUsername(ctx, "johndoe").Return(user, nil)

		response, err := userUsecase.Login(ctx, &model.LoginUserRequest{
			Username: "johndoe",
			Password: "password",
		})
		assert.NoError(t, err)

		userRepository.EXPECT().CountByUsername(ctx, "johndoe").Return(int64(1), nil)

		auth, err := userUsecase.Verify(ctx, &model.VerifyUserRequest{AccessToken: response.AccessToken})
		assert.NoError(t, err)
		assert.Equal(t, user.Username, auth.Username)
	})

	t.Run("failed invalid token", func(t *testing.T) {
		_, err := userUsecase.Verify(ctx, &model.VerifyUserRequest{AccessToken: "wrongToken"})
		assert.Error(t, err)
		assert.ErrorIs(t, exception.ErrUserUnauthorized, err)
	})
}

func createUser(t *testing.T) *domain.User {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	assert.NoError(t, err)

	user := &domain.User{
		ID:        1,
		Name:      "John Doe",
		Username:  "johndoe",
		Password:  string(hashedPassword),
		CreatedAt: 12345,
		UpdatedAt: 12345,
	}

	return user
}
