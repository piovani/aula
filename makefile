start:
	go run main.go

cover:
	go test ./... -coverprofile=covarage.out -covermode=count
	go tool cover -func=covarage.out