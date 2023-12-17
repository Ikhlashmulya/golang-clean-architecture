# golang-clean-architecture-project-structure
This is project structure template following clean architecture and golang standard project layout (not official but commonly used for golang project)

list library/tool has been used in this project
- [fiber](https://github.com/gofiber/fiber) - web framework
- [viper](https://github.com/spf13/viper) - library for configuration
- [gorm](https://github.com/go-gorm/gorm) - ORM library
- [gomock](https://github.com/uber/mock) - mocking framework
- [testify](https://github.com/stretchr/testify)- toolkit for assertions test
- [validator](https://github.com/go-playground/validator) - library for struct and field validation.
- [migrate](https://github.com/golang-migrate/migrate) - database migration

and here's the explanation of the project structure: 

- **cmd** : This directory contains the main entry point of the project. Usually, for the place main application file here, which is used to start and run the application.
- **config**:  This directory holds configuration files for the project.
- **docs**: This directory contains documentation-related files for the project. 
- **internal** : This directory is used to organize internal code of the project. 
- **test** : This directory contains tests (unit tests and integration tests).

other directories :
  - **infrastructure** : This directory contains framework or driver layer in clean architecture
  - **delivery** : This directory contains code related to data delivery, such as HTTP implementation or RPC.
  - **domain**: This directory contains the domain data structure definitions.
  - **exception**: This directory contains code related to error handling.
  - **model** : This directory contains model data structures and response objects used in the project.
  - **mapper** : This directory contains code related to mapping data structures.
  - **repository** : This directory contains code related to data access (data access layer).
  - **usecase** : This directory contains code related to the use case or business logic that governs how data is processed and used in the application.

This project structure looks well-organized and follows many best practices in Go application development. It separates concerns clearly, making the code easy to maintain and extend.

## Install
```
git clone https://github.com/Ikhlashmulya/golang-clean-architecture-project-structure.git
```

```
rm -rf .git/
```

```
cp .env.example .env
```

## Run
```
go run cmd/web/main.go
```
## Testing

### Run Unit Test
```
make test.unit
```

### Run Integration Test
```
make test.integration
```

## Database Migration

### Create
```
make migrate.create name=create_table_users
```

### Up
```
make migrate.up
```

### Down
```
make migrate.down
```

### Force
```
make migrate.force version=20231216100
```
