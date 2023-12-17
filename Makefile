# environment for integration testing
ENV_LOCAL_TEST=\
	APP_NAME="Golang Rest API" \
	APP_PORT=3000 \
	APP_PREFORK=false \
	APP_TIMEOUT=10 \
	DB_USER=root \
	DB_PASSWORD= \
	DB_HOST=localhost \
	DB_PORT=3306 \
	DB_NAME=go-rest-api-test \
	POOL_IDLE=5 \
	POOL_MAX=100 \
	POOL_LIFETIME=3000 \
	LOG_LEVEL=6 \
	JWT_SECRET_KEY=secretkey

test.unit:
	go test ./test/unit -v

test.integration:
	$(ENV_LOCAL_TEST) go test ./test/integration -v

include .env

DATABASE_URL="mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)"

migrate.create:
	migrate create -ext sql -dir db/migrations $(name)

migrate.up:
	migrate -database $(DATABASE_URL) -path db/migrations up

migrate.down:
	migrate -database $(DATABASE_URL) -path db/migrations down

migrate.force:
	migrate -database $(DATABASE_URL) -path db/migrations force $(version)