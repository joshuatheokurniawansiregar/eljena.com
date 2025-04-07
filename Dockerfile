FROM golang:1.23.3-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o my_gin_app

RUN chmod +x my_gin_app

EXPOSE 3000

EXPOSE 8080

CMD ["./my_gin_app"]

