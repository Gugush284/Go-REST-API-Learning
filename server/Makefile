.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: run
run:
	./apiserver.exe

.PHONY: build_run
build_run:
	go build -v ./cmd/apiserver
	./apiserver.exe

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build_run