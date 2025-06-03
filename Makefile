.PHONY: run docker-up docker-build docker-down test coverage test-report swag

run:
	docker-compose up --build

docker-up:
	docker-compose up --build

docker-build:
	docker-compose up --build --no-cache

docker-down:
	docker-compose down -v

docker-clean:
	docker-compose down --volumes --remove-orphans
	docker system prune --all --volumes --force

test:
	go test ./backend/...

coverage:
	go test ./backend/... -cover

test-report:
	go test ./backend/... -coverprofile=coverage.out
	go tool cover -html=coverage.out

swag:
	swag init -g cmd/main.go --dir backend
