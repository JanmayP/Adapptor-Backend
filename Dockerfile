# Build stage
FROM golang:1.23-alpine AS builder

# Set the working directory
WORKDIR /build

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Final stage
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary from builder
COPY --from=builder /build/main .

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["./main"] 