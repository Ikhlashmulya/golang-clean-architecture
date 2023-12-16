package integration

import (
	"testing"

	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/config"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/delivery/http/handler"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/delivery/http/middleware"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/delivery/http/route"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/domain"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/infrastructure"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/repository"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type e2eTestSuite struct {
	suite.Suite
	Config         *viper.Viper
	App            *fiber.App
	DB             *gorm.DB
	Log            *logrus.Logger
	Validate       *validator.Validate
	UserRepository repository.UserRepository
	UserUsecase    usecase.UserUsecase
	UserHandler    *handler.UserHandler
	AuthMiddleware fiber.Handler
	Route          *route.RouteConfig
}

func TestE2eSuite(t *testing.T) {
	suite.Run(t, new(e2eTestSuite))
}

func (s *e2eTestSuite) SetupSuite() {
	s.Config = config.New()
	s.DB = infrastructure.NewGorm(s.Config)
	s.Log = infrastructure.NewLogger(s.Config)
	s.App = infrastructure.NewFiber(s.Config)
	s.Validate = infrastructure.NewValidator(s.Config)
	s.UserRepository = repository.NewUserRepository(s.DB)
	s.UserUsecase = usecase.NewUserUsecase(s.UserRepository, s.Log, s.Validate, s.Config.GetString("JWT_SECRET_KEY"))
	s.UserHandler = handler.NewUserHandler(s.UserUsecase, s.Log)
	s.AuthMiddleware = middleware.NewAuth(s.UserUsecase, s.Log)
	s.Route = route.RegisterRoute(s.App, s.UserHandler, s.AuthMiddleware)
	s.Route.SetupRoute()
}

func (s *e2eTestSuite) SetupTest() {
	s.Require().NoError(s.DB.Migrator().AutoMigrate(&domain.User{}))
}

func (s *e2eTestSuite) TearDownTest() {
	s.Require().NoError(s.DB.Migrator().DropTable("users"))
}
