.PHONY: local
local: prepare-env
	go run github.com/air-verse/air@latest -c .air.toml

.PHONY: prepare-env
prepare-env:
	docker compose up -d
	GOOSE_MIGRATION_DIR=./migrations goose postgres "postgresql://user:password@localhost:5430/likemind?sslmode=disable" up

.PHONY: build
build: build-templ build-tailwind build-app

.PHONY: build-app
build-app:
	go build -o ./bin/likemind ./cmd/main.go

.PHONY: build-templ
build-templ:
	go run github.com/a-h/templ/cmd/templ@latest generate

.PHONY: build-tailwind
build-tailwind:
	npx tailwindcss -i ./website/static/css/src.css -o ./website/static/css/styles.css --minify
