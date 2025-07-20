FROM golang:latest

WORKDIR /app

COPY .env ./

CMD ["tail", "-f", "/dev/null"]
