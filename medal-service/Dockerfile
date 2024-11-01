FROM golang:1.22.5-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o medal-service ./cmd/main.go

EXPOSE 8002

CMD ["./medal-service"]