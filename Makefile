.PHONY: dev-server dev-tailwind dev-templ dev build-server build-tailwind build-templ build launch deploy

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
	@onchange -k -i '**/*.go' '**/*.html' 'views/public/**/*' -- go run ./cmd/server

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
	@go build -o bin/server ./cmd/server/main.go

#-----------------------------------------------------
# DATABASE
#-----------------------------------------------------

# drop:
# 	@go run ./cmd/drop

# seed:
# 	@go run ./cmd/seed

#-----------------------------------------------------
# DEPLOY
#-----------------------------------------------------

launch:
	@fly launch

deploy:
	@fly deploy

.DEFAULT_GOAL := dev