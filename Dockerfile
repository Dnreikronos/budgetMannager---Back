FROM golang:1.22-alpine

WORKDIR /app

COPY  go.mod go.sum ./

RUN go mod download

COPY . . 

ENV DATABASE_URL={DATABASE_URL}

RUN go buil -o main .

CMD ["./main"]