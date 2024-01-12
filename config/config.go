package config

import (
	"os"
)

type postgresqlDBConfig struct {
	ConnectionString string
}

type Config struct {
	DBConfig postgresqlDBConfig
}

func LoadConfig() (*Config, error) {
	// if err := godotenv.Load(); err != nil {
	// 	return nil, errors.New("error to load env")
	// }

	pgsqlConfig := postgresqlDBConfig{ConnectionString: os.Getenv("POSTGRESQL_URL")}

	config := Config{DBConfig: pgsqlConfig}

	return &config, nil
}
