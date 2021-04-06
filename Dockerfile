FROM golang:1.16-alpine AS build
WORKDIR /src
COPY ./ /src
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w"
FROM alpine:latest
WORKDIR /app
COPY --from=build /src/gacha /app
CMD ["/app/gacha"]