# Use an official Golang runtime as a parent image
FROM golang:1.17-alpine

# Set the working directory to /go/src/app
WORKDIR /go/src/app

# Copy the current directory contents into the container at /go/src/app
COPY . .

# Download Go module dependencies
RUN go mod download

# Build the Go binary
RUN go build -o auth

# Expose port 8000 for the service
EXPOSE 8000

# Run the binary
CMD ["./auth"]
