db-up:
	docker-compose  up --remove-orphans --build -d
run:
	go run cmd/api/main.go