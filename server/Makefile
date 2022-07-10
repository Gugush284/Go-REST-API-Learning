.PHONY: buildserver
buildserver:
	go build -v ./cmd/apiserver


.PHONY: runserver
runserver:
	./apiserver.exe



.PHONY: build_run_server
build_run_server:
	go build -v ./cmd/apiserver
	./apiserver.exe



.PHONY: buildclient
buildclient:
	go build -v ./cmd/client



.PHONY: runclient
runserver:
	./client.exe



.PHONY: build_run_client
build_run_server:
	go build -v ./cmd/client
	./client.exe


.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build_run_server