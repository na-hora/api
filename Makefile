include .env
.PHONY: m-generate

dev:
		go run cmd/server/main.go

build:
		go build cmd/server/main.go && ./main

m-generate:
		atlas migrate hash; atlas migrate diff --env gorm $(name)

seed:
		atlas migrate hash; atlas migrate new seed_$(name)

m-apply:
		atlas migrate hash; atlas migrate apply --url ${DB_URL}?sslmode=disable
