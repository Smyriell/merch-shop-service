FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o service ./cmd/service/main.go

FROM alpine:latest
WORKDIR /root/
COPY  --from=builder /app/service . 
EXPOSE 8080
CMD ["./service"]