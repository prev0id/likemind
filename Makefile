.PHONY: local
local:
	npx tailwindcss -i ./website/static/css/src.css -o ./website/static/css/styles.css --watch&
	go run github.com/air-verse/air@latest -c .air.toml
