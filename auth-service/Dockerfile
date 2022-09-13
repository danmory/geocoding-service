FROM golang:1.19-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /auth-server github.com/danmory/geocoding-service/auth-service/cmd/main

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /auth-server /auth-server

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/auth-server"]
