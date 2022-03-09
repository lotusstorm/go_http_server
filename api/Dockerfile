FROM golang:1.17-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go install github.com/cespare/reflex@latest

EXPOSE 5000
CMD exec reflex -r '(\.go$|go\.mod)' go run main.go --start-service
