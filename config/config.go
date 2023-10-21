package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type PostgresqlDBConfig struct {
	ConnectionString string
}

type Config struct {
	DBConfig PostgresqlDBConfig
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("could not load .env")
	}

	pgsqlConfig := PostgresqlDBConfig{ConnectionString: os.Getenv("POSTGRESQL_URL")}

	config := Config{DBConfig: pgsqlConfig}

	return &config, nil
}
