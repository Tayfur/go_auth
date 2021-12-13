package db

import (
	"fmt"
	"go_auth/models"
	"go_auth/pkg/config"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb() (*gorm.DB, error) {
	var dsn string
	if len(config.Config("DBUrl")) == 0 {
		dsn = fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
			config.Config("DB_HOST"),
			config.PortConfig("DB_PORT"),
			config.Config("POSTGRES_USER"),
			config.Config("POSTGRES_DB"),
			config.Config("POSTGRES_PASSWORD"))
	} else {
		dsn = config.Config("DBUrl")
	}
	database, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Error(err)
		panic(err)
	}
	err = database.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}
	return database, nil
}
