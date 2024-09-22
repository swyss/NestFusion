# Stage 1: Build the Go binary
FROM golang:1.22 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go binary inside the container
RUN go build -o /app/main

# Stage 2: Run the Go application
FROM golang:1.22-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the built Go binary from the builder stage
COPY --from=builder /app/main .

# Expose the port on which the app will run
EXPOSE 8080

# Command to run the application
CMD ["./main"]
