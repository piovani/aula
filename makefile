setup:
	cp .env.example .env
	docker-compose up -d
	docker exec aula-app go run main.go api

start:
	go run main.go

cover:
	go test ./... -coverprofile=covarage.out -covermode=count
	go tool cover -func=covarage.out