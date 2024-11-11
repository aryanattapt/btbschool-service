# Stage 1: Build the Go binary
FROM golang:1.22.4 AS builder

# Set the current working directory inside the container
WORKDIR /app/btbschool-service

# Copy go.mod and go.sum to leverage Docker cache for dependencies
COPY go.mod go.sum ./

# Download and cache dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go binary
RUN go build -o btbschool-service

# Stage 2: Create the final image with only the binary
FROM alpine:latest

# Install necessary dependencies (e.g., for running Go binary in Alpine)
RUN apk add --no-cache libc6-compat

# Set the working directory
WORKDIR /app/btbschool-service

# Copy the binary from the builder stage
COPY --from=builder /app/btbschool-service /app/btbschool-service

# Expose the port the app will run on
EXPOSE 30001

# Set the entry point to run the Go binary
ENTRYPOINT ["/app/btbschool-service/btbschool-service"]