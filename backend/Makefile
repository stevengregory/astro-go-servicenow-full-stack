.PHONY: all build clean run

all: clean build run

build:
	@go build -o bin/service cmd/service/*.go

clean:
	@rm -rf bin

run:
	@bin/service