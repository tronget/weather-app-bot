# Build stage
FROM golang:1.25-alpine AS builder

# Set working directory
WORKDIR /app

# Install git (needed for go mod download)
RUN apk add --no-cache git

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates tzdata

# Set timezone
ENV TZ=UTC

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Copy the entire internal directory to preserve structure
COPY --from=builder /app/internal ./internal

# Copy .env file (will be overridden by docker-compose)
COPY --from=builder /app/.env .

# Expose port (if you plan to add a health check endpoint later)
EXPOSE 8080

# Run the binary
CMD ["./main"]
