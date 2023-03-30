# Defaults
PORT?=3000

run:
	@go run main.go --port=${PORT}

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down