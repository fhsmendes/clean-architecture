FROM golang:latest

WORKDIR /app

# CMD ["sh", "-c", "go run main.go wire_gen.go"]

# CMD ["go", "run", "main.go wire_gen.go"]

COPY cmd/ordersystem/.env ./

CMD ["tail", "-f", "/dev/null"]
