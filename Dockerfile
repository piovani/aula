FROM golang:1.22.3-alpine3.20 as builder

WORKDIR /src

COPY . .

RUN go mod download -x

RUN GOOS=linux GORASCH=amd64 go build -o server

FROM alpine:3.20

EXPOSE 8080

WORKDIR /app

COPY --from=builder /src/server .

ENTRYPOINT ["./server", "api"]