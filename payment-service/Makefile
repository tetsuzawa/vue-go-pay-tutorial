BIN_NAME=payment-service


.PHONY: build
build:
	mkdir -p ./bin/
	go build -o bin/$(BIN_NAME) .


.PHONY: install
install:
	go install

.PHONY: test
test:
	go test -cover -v -race
