package app

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDb() *gorm.DB {
	env := os.Getenv("ENV")
	dsn := os.Getenv("DB_DSN")
	if env == "test" {
		dsn = os.Getenv("DB_DSN_TEST")
	}

	var err error

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
