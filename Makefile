.PHONY: docker-up docker-build docker-down test coverage test-report

docker-up:
	docker-compose up --build

docker-build:
	docker-compose up --build --no-cache

docker-down:
	docker-compose down -v

test:
	go test ./backend/...

coverage:
	go test ./backend/... -cover

test-report:
	go test ./backend/... -coverprofile=coverage.out
	go tool cover -html=coverage.out
