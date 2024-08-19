include .env

.ONESHELL:

build:
	@go build -o ./.bin/pokedex	./cmd/pokedex

run:build
	@./.bin/pokedex

up:
	cd ./sql/migrations;
	goose postgres $(DATABASE_URL) up

down:
	cd ./sql/migrations;
	goose postgres $(DATABASE_URL) down