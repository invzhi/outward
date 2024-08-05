.PHONY: all
all: build lint

.PHONY: build
build:
	go build -o outward cmd/outward/main.go

.PHONY: run
run:
	go run cmd/outward/main.go

.PHONY: lint
lint:
	golangci-lint run --fix

.PHONY: sqlc
sqlc:
	sqlc generate

.PHONY: proto
proto:
	buf dep update && buf generate

.PHONY: clean
clean:
	rm -f outward
