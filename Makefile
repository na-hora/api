include .env
.PHONY: m-generate

r-dev:
		make m-apply; air

dev:
		make m-apply; go build -o ./tmp/main cmd/server/main.go; ./tmp/main

build:
		go build -o ./tmp/main cmd/server/main.go

m-generate:
		atlas migrate hash; atlas migrate diff --env gorm $(name)

seed:
		atlas migrate hash; atlas migrate new seed_$(name)

m-apply:
		atlas migrate hash; atlas migrate apply --url ${DB_URL}

wire:
		cd internal/injector && wire && cd ../..
