# golang-clean-architecture-project-structure
This is project structure template following clean architecture and golang standard project layout (not official but commonly used for golang project)

list library/tool has been used in this project
- [fiber](https://github.com/gofiber/fiber) - web framework
- [goDotEnv](https://github.com/joho/godotenv) - library for loads environment variables from .env files
- [sqlmock](https://github.com/DATA-DOG/go-sqlmock) - library for test database interactions
- [mysql](https://github.com/go-sql-driver/mysql) - driver mysql
- [gomock](https://github.com/uber/mock) - mocking framework
- [swaggo](https://github.com/swaggo/swag) - tool for generate RESTful API documentation
- [testify](https://github.com/stretchr/testify)- toolkit for assertions test
- [validator](https://github.com/go-playground/validator) - library for struct and field validation.
- [migrate](https://github.com/golang-migrate/migrate) - database migration

and here's the explanation of the project structure: 

- **cmd** : This directory contains the main entry point of the project. Usually, for the place main application file here, which is used to start and run the application.
- **config**:  This directory holds configuration files for the project.
- **docs**: This directory contains documentation-related files for the project. 
- **internal** : This directory is used to organize internal code of the project. 

other directories :

  - **infrastructure** : This directory contains framework or driver layer in clean architecture
  - **delivery** : This directory contains code related to data delivery, such as HTTP implementation or RPC.
  - **domain**: This directory contains the domain data structure definitions.
  - **exception**: This directory contains code related to error handling.
  - **injector** : This directory contains `wire.go` and `wire_gen.go` files related to Dependency Injection (DI).
  - **mapper** : This directory contains code related to mapping data structures.
  - **mock** : This directory contains code for mock objects used in testing.
  - **model** : This directory contains model data structures and response objects used in the project.
  - **repository** : This directory contains code related to data access (data access layer).
  - **test** : This directory contains tests (unit tests and integration tests).
  - **usecase** : This directory contains code related to the use case or business logic that governs how data is processed and used in the application.

This project structure looks well-organized and follows many best practices in Go application development. It separates concerns clearly, making the code easy to maintain and extend.