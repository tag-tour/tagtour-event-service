.SILENT:

all: run

build:
	go mod -C src tidy
	go build -C src -o ../bin/service ./cmd/main.go

run: build
	./bin/service