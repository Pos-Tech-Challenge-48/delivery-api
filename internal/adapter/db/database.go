package db

import (
	"log"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	dbURL := "postgres://postech:postech@db:5432/postech?sslmode=disable"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
