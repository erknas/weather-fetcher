city ?=
port ?= :6000

build:
	@go build -o bin/weatherfetcher

run: build
	@./bin/weatherfetcher -city=$(city) -port=$(port)

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/service.proto

.PHONY: proto