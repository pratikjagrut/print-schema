# Use the official Go image as the base image
FROM golang:1.16-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download and cache Go module dependencies
RUN go mod download

# Copy the application source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o app

# Use a minimal Alpine Linux image as the base image for the final container
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the previous stage
COPY --from=builder /app/app .

# Set the environment variable for the database URL
ENV DB_URL="your_database_url_here"

# Expose the port that the application listens on
EXPOSE 8080

# Run the Go application
CMD ["./app"]
