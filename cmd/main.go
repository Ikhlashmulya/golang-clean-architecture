package main

import (
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/config"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/exception"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/injector"

	_ "github.com/Ikhlashmulya/golang-clean-architecture-project-structure/docs"
)

//	@title			golang-clean-architecture-project-structure
//	@version		1.0
//	@description	This is a sample project following clean architecture.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Ikhlash Mulyanurahman
//	@contact.url	https://www.ikhlashmulya.github.io/

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host			localhost:8080
// @BasePath		/api
func main() {
	configuration := config.NewConfig()
	app := injector.InitializedApp(configuration)

	err := app.Listen(":8080")
	exception.PanicIfError(err)
}
