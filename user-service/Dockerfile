# Build stage
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o user-service cmd/main.go

# Run stage
FROM alpine:3.18

WORKDIR /root/

COPY --from=builder /app/user-service .
COPY .env .

EXPOSE 3001

CMD ["./user-service"]
