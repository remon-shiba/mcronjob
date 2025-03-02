# Use the official Golang image as the base image
FROM golang:1.24 as builder

# Set the working directory
WORKDIR /app

# Copy the Go modules manifests and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the rest of the application source code
COPY . .

# Build the application
RUN go build -o main .

# Use a minimal image for deployment
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose the port the application runs on
EXPOSE 8080

# Command to run the application
CMD ["./main"]
