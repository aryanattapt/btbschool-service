# Stage 1: Build the Go binary
FROM golang:1.22.4 AS builder

# Set the current working directory inside the container
WORKDIR /app/btbschool-service

# Set environment variables
ENV DOMAIN=".aryanattapt.my.id" \
    PORT="30001" \
    ENV="PRODUCTION" \
    POSTGRESQL_HOST="localhost" \
    POSTGRESQL_PORT="20002" \
    POSTGRESQL_USERNAME="postgresql" \
    POSTGRESQL_PASSWORD="postgresql" \
    MONGODB_URL="mongodb://mongodb:mongodb@localhost:20001" \
    REDIS_URL="" \
    REDIS_PASSWORD="" \
    ASSET_PATH="./assets" \
    UPLOAD_PATH="./uploads" \
    JWT_SIGNATURE_KEY="SIGNATURE" \
    MAIL_SMTP_HOST="" \
    MAIL_SMTP_PORT="" \
    MAIL_SENDER_NAME="" \
    MAIL_AUTH_EMAIL="" \
    MAIL_AUTH_PASSWORD="" \
    API_BASE_URL="localhost:30001" \
    RECAPTCHA_SITE_KEY="6Lc5SikqAAAAAIury1pPE5QsX1ilLuyVL8MsXdd_" \
    RECAPTCHA_SECRET_KEY="6Lc5SikqAAAAAE3pIs6vKTMZGZgqtSj43E1bTUwY"

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
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/btbschool-service /app/

# Expose the port the app will run on
EXPOSE 30001

# Set the entry point to run the Go binary
ENTRYPOINT ["/app/btbschool-service"]