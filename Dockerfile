# Builder stage
FROM golang:alpine AS builder

WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application statically
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o log-collector cmd/main.go

# Final minimal stage
FROM alpine:latest

# Install certificates for HTTPS requests if needed
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/log-collector .

# Create the folder where the host logs will be mounted
RUN mkdir -p /var/log/collector

# Expose API port
EXPOSE 8000

# Run the binary
ENTRYPOINT ["./log-collector"]
