FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o go_server github.com/danmory/geocoding-service/data-service/cmd/main

EXPOSE 8081

ENTRYPOINT ["./go_server"]
