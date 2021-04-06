FROM golang:1.16-alpine AS build
WORKDIR /app
COPY go.mod go.sum /app/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go mod download
COPY ./ /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w"
CMD ["/app/gacha"]