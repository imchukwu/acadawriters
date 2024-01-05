# Dockerfile

# Build phase
FROM golang:latest AS builder
WORKDIR /app

# Copy only necessary files to avoid unnecessary cache invalidation
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o out ./cmd/main

# Final stage
FROM alpine:latest
WORKDIR /app

# Copy only the compiled binary from the builder stage
COPY --from=builder /app/out .

# Command to run the application
CMD ["./out"]
