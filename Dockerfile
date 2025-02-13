# 🔹 Stage 1: Build the Go binary
FROM golang:1.23.5 AS builder

# Set working directory
WORKDIR /app

# Copy Go modules and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application (output as `app`)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/api/main.go

# 🔹 Stage 2: Create a minimal runtime image
FROM alpine:latest

# Install PostgreSQL client (for health checks/debugging)
RUN apk add --no-cache postgresql-client

# Set working directory
WORKDIR /root/

# Copy the compiled binary from builder stage
COPY --from=builder /app/app .

# Expose the port your Go app runs on
EXPOSE 8080

# Run the application
CMD ["./app"]
