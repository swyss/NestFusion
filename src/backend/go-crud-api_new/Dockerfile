# Use the official Golang image as the base
FROM golang:1.22

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Download and install dependencies
RUN go get -d -v ./...

# Build the Go application
RUN go build -o api .

# Expose the application port
EXPOSE 8000

# Set the command to run the application
CMD ["./api"]
