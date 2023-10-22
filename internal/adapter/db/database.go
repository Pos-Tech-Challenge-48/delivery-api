package db

import (
	"log"

	"github.com/Pos-Tech-Challenge-48/delivery-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(cfg *config.Config) *gorm.DB {
	dbURL := cfg.DBConfig.ConnectionString

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
