package infrastructure 

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/config"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/exception"
	_ "github.com/go-sql-driver/mysql"
)

func NewDB(configuration *config.Config) *sql.DB {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			configuration.Get("DB_USERNAME"),
			configuration.Get("DB_PASSWORD"),
			configuration.Get("DB_HOST"),
			configuration.Get("DB_PORT"),
			configuration.Get("DB_NAME"),
		),
	)
	exception.PanicIfError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(50)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(time.Hour)

	return db
}
