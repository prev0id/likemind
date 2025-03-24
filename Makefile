.PHONY: prepare-env
prepare-env: docker-up migrate create-bucket

.PHONY: docker-down
docker-down:
	docker compose down -v

.PHONY: docker-up
docker-up:
	docker compose up -d

.PHONY: migrate
migrate:
	sleep 5s
	GOOSE_MIGRATION_DIR=./migrations goose postgres "postgresql://user:password@localhost:5432/likemind?sslmode=disable" up

.PHONY: create-bucket
create-bucket:
	docker-compose exec minio mc mb minio/my-bucket

.PHONY: build
build: build-templ build-tailwind build-app

.PHONY: local
local:
	go run github.com/air-verse/air@v1.61.5 -c .air.toml

.PHONY: run
run: build-templ build-tailwind
	go run ./cmd/main.go

.PHONY: build-app
build-app:
	go build -o ./bin/likemind ./cmd/main.go

.PHONY: build-templ
build-templ:
	go run github.com/a-h/templ/cmd/templ@latest generate

.PHONY: build-tailwind
build-tailwind:
	npx tailwindcss -i ./website/static/css/src.css -o ./website/static/css/styles.css --minify
