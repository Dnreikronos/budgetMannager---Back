FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV DATABASE_URL={$DATABASE_URL}

RUN go build -o cmd/main ./cmd


EXPOSE 9090

CMD ["./main"]
