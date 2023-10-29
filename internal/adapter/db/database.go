package db

import (
	"database/sql"
	"log"

	"github.com/Pos-Tech-Challenge-48/delivery-api/config"
)

func NewDatabase(config *config.Config) *sql.DB {

	connectionString := config.DBConfig.ConnectionString

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
