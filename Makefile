run:
	@go run main.go

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down