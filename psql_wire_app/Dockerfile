# Stage 1: The builder
# Use an official Go image to build the application
FROM golang:1.24-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app, creating a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /main .

# Stage 2: The final image
# Use a minimal Alpine image for a small footprint
FROM alpine:latest

# Copy the built binary from the builder stage
COPY --from=builder /main /main

# Expose the default PostgreSQL port.
# Your Go app must listen on this port inside the container.
EXPOSE 5432

# Command to run the executable
CMD ["/main"]