BIN_NAME=backend-api

mkdir -p bin

.PHONY: build
build:
	go build -o bin/$(BIN_NAME) .


.PHONY: install
install:
	go install

.PHONY: test
test:
	go test -cover -v -race
