package db

import (
	"log"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	dbURL := "postgres://pg:pass@localhost:5432/user"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
