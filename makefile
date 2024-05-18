setup:
	cp .env.example .env
	docker-compose up -d
	docker exec aula-app go run main.go api

start:
	air api 

cover:
	go test ./... -coverprofile=covarage.out -covermode=count
	go tool cover -func=covarage.out

build:
	GOOS=linux GOARCH=amd64 go build -o service