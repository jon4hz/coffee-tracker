FROM golang:1.17-alpine

WORKDIR /app

RUN apk update --no-cache && \
    apk add gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build cmd/coffee-tracker/main.go