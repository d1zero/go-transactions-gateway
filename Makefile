.SILENT:

gen-docs:
	swag init -g internal/app/app.go

run:
	go run cmd/app/main.go

test:
	go test ./...

test-verbose:
	go test -v ./...

lint:
	golangci-lint run

lint-verbose:
	golangci-lint run -v

db-up:
	docker compose up -d

db-down:
	docker compose down -v --rmi "local"