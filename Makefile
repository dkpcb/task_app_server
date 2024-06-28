# Docker Compose でサービスをビルドしてバックグラウンドで起動
dev/run/import:
	docker-compose up -d --build
	docker-compose exec server go run main.go