include .env

local-install:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

setup-dev:
	migrate -database ${POSTGRESQL_URL} -path db/migrations up
