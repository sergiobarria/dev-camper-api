run:
	@go run ./cmd/api

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down