.PHONY: build run

build:
	go build -tags netgo -ldflags '-s -w' -o app ./cmd/server

run:
	docker-compose up
