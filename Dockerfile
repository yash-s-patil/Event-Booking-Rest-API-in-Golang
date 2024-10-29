# Build stage
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o api-server .

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/api-server .
EXPOSE 8080
CMD ["./api-server"]
