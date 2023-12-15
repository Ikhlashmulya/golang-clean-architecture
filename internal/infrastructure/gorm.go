package infrastructure

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGorm(config *viper.Viper) *gorm.DB {
	user := config.GetString("database.user")
	password := config.GetString("database.password")
	host := config.GetString("database.host")
	dbname := config.GetString("database.name")
	port := config.GetInt("database.port")
	idleConns := config.GetInt("database.pool.idle")
	maxConns := config.GetInt("database.pool.max")
	lifetime := config.GetInt("database.pool.lifetime")

	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbname,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(fmt.Errorf("error connecting database : %+v", err.Error()))
	}

	connection, err := db.DB()
	if err != nil {
		panic(err)
	}

	connection.SetMaxIdleConns(idleConns)
	connection.SetMaxOpenConns(maxConns)
	connection.SetConnMaxLifetime(time.Duration(lifetime))

	return db
}
