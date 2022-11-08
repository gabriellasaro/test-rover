FROM golang:1.19-alpine

WORKDIR /app

COPY . .

RUN go test ./... && \
    go build -o /app/testrover ./cmd && \
    ./testrover data/input.txt
