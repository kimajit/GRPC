# Use the official Golang image as the base image for building the Go application
FROM golang:1.21.4 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download all dependencies specified in go.mod and go.sum
RUN go mod download

# Copy the entire project to the working directory
COPY . .

# Build the Go application with static linking
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o userdata .

# Use a minimal image for the runtime
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Install necessary packages (e.g., for debug)
RUN apk add --no-cache ca-certificates

# Copy the built Go binary from the builder stage to the current stage
COPY --from=builder /app/userdata .

# Verify the contents and permissions of /app directory
RUN ls -la /app

# Expose the port on which the gRPC server will run
EXPOSE 50051

# Command to run the application
CMD ["./userdata"]
