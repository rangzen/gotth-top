.PHONY: dev-server dev-tailwind dev-templ dev build-server build-tailwind build-templ build

#-----------------------------------------------------
# DEV
#-----------------------------------------------------

dev:
	@make -j dev-tailwind dev-templ dev-server

dev-tailwind:
	@make ARGS="--watch" build-tailwind

dev-templ:
	@templ generate --watch

dev-server:
	@onchange -k -i '**/*.go' '**/*.html' 'views/public/**/*' -- go run ./cmd/gotth-top

#-----------------------------------------------------
# BUILD
#-----------------------------------------------------

build:
	@make build-tailwind build-templ build-server

build-tailwind:
	@npx tailwindcss -m -i ./tailwind.css -o ./views/public/styles.css $(ARGS)

build-templ:
	@templ generate

build-server:
	@go build -ldflags="-s -w" -o bin/gotth-top ./cmd/gotth-top/main.go

.DEFAULT_GOAL := build
