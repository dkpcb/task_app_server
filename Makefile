app:
	docker-compose up -d --build
	docker-compose exec server go run main.go