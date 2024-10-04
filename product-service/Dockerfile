# Build stage
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o product-service cmd/main.go

# Run stage
FROM alpine:3.18

WORKDIR /root/

COPY --from=builder /app/product-service .

EXPOSE 3000

CMD ["./product-service"]
