PROJECT_NAME = survey

.PHONY: golangci
golangci:
	golangci-lint run

.PHONY: gofmt
gofmt:
	gofmt -s -w -l .

.PHONY: goimport
goimport:
	goimports -local ./ -w .

.PHONY: lint
lint: gofmt goimport

.PHONY: test
test:
	go test $(shell go list ./... | grep -v /storage) -v -covermode=count

.PHONY: build
build:
	go build -v -o ./cmd/$(PROJECT_NAME) ./cmd/main.go

.PHONY: run
run:
	go run cmd/main.go

.PHONY: docker-run
docker-run:
	docker-compose up -d