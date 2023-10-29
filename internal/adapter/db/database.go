package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Pos-Tech-Challenge-48/delivery-api/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // postgres driver
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func NewDatabase(config *config.Config) *sql.DB {

	connectionString := config.DBConfig.ConnectionString

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatalln(err)
	}

	err = runMigrations(connectionString)

	if err != nil {
		log.Fatal(err)
	}
	return db
}

func runMigrations(connString string) error {
	migrationsPath := "db/migrations"

	fmt.Println("db: creating migrations")
	m, err := migrate.New("file://"+migrationsPath, connString)
	if err != nil {
		return fmt.Errorf("creating migrate instance: %w", err)
	}

	if err = m.Up(); err != nil {
		return fmt.Errorf("running up migrations: %w", err)
	}

	serr, err := m.Close()
	if serr != nil {
		return fmt.Errorf("closing the source: %w", serr)
	}

	if err != nil {
		return fmt.Errorf("closing postgres connection: %w", err)
	}

	return nil
}
