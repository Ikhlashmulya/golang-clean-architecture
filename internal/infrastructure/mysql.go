package infrastructure

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/exception"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func NewDB(configuration *viper.Viper) *sql.DB {
	username := configuration.GetString("database.username")
	password := configuration.GetString("database.password")
	host := configuration.GetString("database.host")
	dbname := configuration.GetString("database.name")
	port := configuration.GetString("database.port")

	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname),
	)
	exception.PanicIfError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(50)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(time.Hour)

	return db
}

func NewDBTesting(configuration *viper.Viper) *sql.DB {
	username := configuration.GetString("database_test.username")
	password := configuration.GetString("database_test.password")
	host := configuration.GetString("database_test.host")
	dbname := configuration.GetString("database_test.name")
	port := configuration.GetString("database_test.port")

	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname),
	)
	exception.PanicIfError(err)

	return db
}
