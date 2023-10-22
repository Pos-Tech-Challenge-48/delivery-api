include .env

local-install:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

setup-dev:
	migrate -database ${POSTGRESQL_URL_PUBLIC} -path db/migrations up


migrate-down:
	migrate -database ${POSTGRESQL_URL_PUBLIC} -path db/migrations down
