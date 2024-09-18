# Dockerfile for the Go app starter

# Use the official Golang image as the base
FROM golang:1.22

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the application port
EXPOSE 8000

# Set the default command for Docker, to either reset Docker or start the app
CMD ["./main"]
