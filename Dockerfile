# Use the official Go image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download and install the Go dependencies
RUN go mod download

# Copy the rest of the application source code
COPY *.go .

# Build the Go application
RUN go build -o /main

EXPOSE 8080

# Set the entry point for the container
CMD ["/main"]