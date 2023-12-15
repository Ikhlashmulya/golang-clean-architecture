package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/config"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/delivery/http"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/infrastructure"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/repository"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/usecase"
	"github.com/go-playground/validator/v10"
)

func main() {
	config := config.New()

	app := infrastructure.NewFiber(config)
	port := config.Get("app.port")

	db := infrastructure.NewGorm(config)
	logger := infrastructure.NewLogger(config)
	validate := validator.New()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, logger, validate)
	userHandler := http.NewUserHandler(userUsecase, logger)

	app.Post("/api/login", userHandler.Login)
	app.Post("/api/register", userHandler.Register)


	go func() {
		if err := app.Listen(fmt.Sprintf(":%v", port)); err != nil {
			panic(fmt.Errorf("error running app : %+v", err.Error()))
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the

	<-c // This blocks the main thread until an interrupt is received

	// Your cleanup tasks go here
	_ = app.Shutdown()

	fmt.Println("App was successful shutdown.")
}
