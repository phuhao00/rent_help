# Multi-stage Dockerfile for Go backend

# Development stage
FROM golang:1.21-alpine AS development

WORKDIR /app

# Set Go proxy for faster downloads in China
ENV GOPROXY=https://goproxy.cn,direct
ENV GOSUMDB=sum.golang.org

# Install Air for hot reloading in development
RUN go install github.com/cosmtrek/air@v1.49.0

# Install curl for health checks
RUN apk add --no-cache curl

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Create uploads directory
RUN mkdir -p /app/uploads

# Expose port
EXPOSE 8080

# Default command for development (will be overridden by docker-compose)
CMD ["air", "-c", ".air.toml"]

# Build stage for production
FROM golang:1.21-alpine AS base

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server

# Production stage
FROM alpine:latest AS production

RUN apk --no-cache add ca-certificates curl

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Create uploads directory
RUN mkdir -p /root/uploads

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./main"]

# Run the binary
CMD ["./main"]
