package usecase

import (
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/domain"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/exception"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/model"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/model/mapper"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type UserUsecase interface {
	Register(ctx context.Context, request *model.RegisterUserRequest) (*model.UserResponse, error)
	Login(ctx context.Context, request *model.LoginUserRequest) (*model.TokenResponse, error)
	Update(ctx context.Context, request *model.UpdateUserRequest) (*model.UserResponse, error)
	Current(ctx context.Context, request *model.GetUserRequest) (*model.UserResponse, error)
	Verify(ctx context.Context, request *model.VerifyUserRequest) (*model.Auth, error)
}

type UserUsecaseImpl struct {
	UserRepository repository.UserRepository
	Logger         *logrus.Logger
	Validate       *validator.Validate
	Config         *viper.Viper
}

func NewUserUsecase(userRepo repository.UserRepository, log *logrus.Logger,
	validate *validator.Validate, config *viper.Viper) UserUsecase {
	return &UserUsecaseImpl{
		UserRepository: userRepo,
		Logger:         log,
		Validate:       validate,
		Config:         config,
	}
}

func (uc *UserUsecaseImpl) Register(ctx context.Context, request *model.RegisterUserRequest) (*model.UserResponse, error) {
	if err := uc.Validate.Struct(request); err != nil {
		uc.Logger.WithError(err).Error("failed validating request body")
		return nil, err
	}

	countUser, err := uc.UserRepository.CountByUsername(ctx, request.Username)
	if err != nil {
		uc.Logger.WithError(err).Error("failed count user by username")
		return nil, exception.ErrInternalServerError
	}

	if countUser > 0 {
		uc.Logger.Warn("user already exists")
		return nil, exception.ErrUserAlreadyExist
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		uc.Logger.WithError(err).Error("failed hashing password")
		return nil, exception.ErrInternalServerError
	}

	user := new(domain.User)
	user.Name = request.Name
	user.Username = request.Username
	user.Password = string(hashedPassword)

	if err := uc.UserRepository.Create(ctx, user); err != nil {
		uc.Logger.WithError(err).Error("failed create user to database")
		return nil, exception.ErrInternalServerError
	}

	return mapper.ToUserResponse(user), nil
}

func (uc *UserUsecaseImpl) Login(ctx context.Context, request *model.LoginUserRequest) (*model.TokenResponse, error) {
	if err := uc.Validate.Struct(request); err != nil {
		uc.Logger.WithError(err).Error("failed validating request body")
		return nil, err
	}

	user, err := uc.UserRepository.FindByUsername(ctx, request.Username)
	if err != nil {
		uc.Logger.WithError(err).Error("failed find user by username")
		return nil, exception.ErrUserNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		uc.Logger.WithError(err).Error("failed to compare hashedPassword and password")
		return nil, exception.ErrUserPasswordNotMatch
	}

	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(2 * time.Hour).Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(uc.Config.GetString("JWT_SECRET_KEY")))
	if err != nil {
		uc.Logger.WithError(err).Error("failed sign token")
		return nil, exception.ErrInternalServerError
	}

	tokenResponse := &model.TokenResponse{
		AccessToken: token,
		TokenType:   "Bearer",
	}

	return tokenResponse, nil
}

func (uc *UserUsecaseImpl) Update(ctx context.Context, request *model.UpdateUserRequest) (*model.UserResponse, error) {
	if err := uc.Validate.Struct(request); err != nil {
		uc.Logger.WithError(err).Error("failed validating request body")
		return nil, err
	}

	user, err := uc.UserRepository.FindByUsername(ctx, request.Username)
	if err != nil {
		uc.Logger.WithError(err).Error("failed find user by username")
		return nil, exception.ErrUserNotFound
	}

	if request.Name != "" {
		user.Name = request.Name
	}

	if request.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			uc.Logger.WithError(err).Error("failed hashing password")
			return nil, exception.ErrInternalServerError
		}

		user.Password = string(hashedPassword)
	}

	if err := uc.UserRepository.Update(ctx, user); err != nil {
		uc.Logger.WithError(err).Error("failed update user to database")
		return nil, exception.ErrInternalServerError
	}

	return mapper.ToUserResponse(user), nil
}

func (uc *UserUsecaseImpl) Current(ctx context.Context, request *model.GetUserRequest) (*model.UserResponse, error) {
	if err := uc.Validate.Struct(request); err != nil {
		uc.Logger.WithError(err).Error("failed validating request body")
		return nil, err
	}

	user, err := uc.UserRepository.FindByUsername(ctx, request.Username)
	if err != nil {
		uc.Logger.WithError(err).Error("failed find user by username")
		return nil, exception.ErrUserNotFound
	}

	return mapper.ToUserResponse(user), nil
}

func (uc *UserUsecaseImpl) Verify(ctx context.Context, request *model.VerifyUserRequest) (*model.Auth, error) {
	token, err := jwt.Parse(request.AccessToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(uc.Config.GetString("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		uc.Logger.WithError(err).Error("user unauthorized")
		return nil, exception.ErrUserUnauthorized
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, exception.ErrUserUnauthorized
	}

	countUser, err := uc.UserRepository.CountByUsername(ctx, claims["username"].(string))
	if err != nil {
		uc.Logger.WithError(err).Error("failed count user by username")
		return nil, exception.ErrInternalServerError
	}

	if countUser == 0 {
		return nil, exception.ErrUserUnauthorized
	}

	return &model.Auth{
		Username: claims["username"].(string),
		ID:       uint(claims["id"].(float64)),
	}, nil
}
