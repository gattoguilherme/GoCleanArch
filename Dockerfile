# Stage 1: Build the application
FROM golang:1.24-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies first
# This leverages Docker's layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application as a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/server

# Stage 2: Create the final, minimal image
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/server .

# Copy the configs directory needed for the application to run
COPY configs/ ./configs/

# Expose the port the application runs on
EXPOSE 8090

# Command to run the application
CMD ["./server", "-config", "./configs/config.yaml"]
