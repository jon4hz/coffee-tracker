run:
	go run cmd/coffee-tracker/main.go

docker.run:
	docker-compose -f docker-compose.dev.yml up --build