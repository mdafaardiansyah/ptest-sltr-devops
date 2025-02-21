# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod dan sum files
COPY go.* ./
RUN go mod download

# Copy source code
COPY . .

# Build aplikasi
RUN CGO_ENABLED=0 GOOS=linux go build -o /welcome-app ./cmd/main.go

# Final stage
FROM alpine:3.18

WORKDIR /app

# Copy binary dari build stage
COPY --from=builder /welcome-app .

# Expose port
EXPOSE 5000

# Run aplikasi
CMD ["./welcome-app"]