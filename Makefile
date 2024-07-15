build:
	@go build -o ./.bin/pokedex	./cmd/pokedex

run:build
	@./.bin/pokedex