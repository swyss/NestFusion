# Use the official Golang image as the base
FROM golang:1.22
#RUN go get -u github.com/cosmtrek/air
# Set the working directory inside the container
WORKDIR /app

# Install Swag CLI for generating Swagger documentation
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Copy the source code into the container
COPY . .

# Download and install dependencies
RUN go get -d -v ./...

# Generate Swagger documentation
#RUN swag init

# Build the Go application
RUN go build -o api .

# Expose the application port
EXPOSE 8000

# Set the command to run the application
CMD ["./api"]
