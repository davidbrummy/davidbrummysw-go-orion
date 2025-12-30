FROM golang:1.25.5-alpine

RUN apk update \
    && apk add --no-cache \
    build-base

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build cmd/main.go

ENTRYPOINT ["/app/main"]