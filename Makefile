build:
	go build -o bin/weatherfetcher

run: build
	./bin/weatherfetcher