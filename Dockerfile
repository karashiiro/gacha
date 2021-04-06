FROM golang:1.16-alpine
WORKDIR /app
COPY ./ /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w"
FROM alpine:latest
CMD ["/app/gacha"]