package db

import (
	"log"

	"github.com/ffelipelimao/Pos-Tech-Challenge-48/delivery-api/config"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

func NewDatabase(config *config.Config) *gorm.DB {

	db, err := gorm.Open(postgres.Open(config.DBConfig.ConnectionString), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
