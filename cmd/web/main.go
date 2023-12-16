package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/config"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/delivery/http/handler"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/delivery/http/middleware"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/delivery/http/route"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/infrastructure"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/repository"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/usecase"
	"github.com/go-playground/validator/v10"
)

func main() {
	config := config.New()

	app := infrastructure.NewFiber(config)
	port := config.Get("APP_PORT")

	db := infrastructure.NewGorm(config)
	logger := infrastructure.NewLogger(config)
	validate := validator.New()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, logger, validate, config.GetString("JWT_SECRET_KEY"))
	userHandler := handler.NewUserHandler(userUsecase, logger)

	authMiddleware := middleware.NewAuth(userUsecase, logger)

	route := route.RegisterRoute(app.Group("/api"), userHandler, authMiddleware)
	route.SetupRoute()

	go func() {
		if err := app.Listen(fmt.Sprintf(":%v", port)); err != nil {
			panic(fmt.Errorf("error running app : %+v", err.Error()))
		}
	}()

	ch := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	<-ch // This blocks the main thread until an interrupt is received

	// Your cleanup tasks go here
	_ = app.Shutdown()

	fmt.Println("App was successful shutdown.")
}
