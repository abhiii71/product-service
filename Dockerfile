# Use the official Go image as a base image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules and vendor files
COPY go.mod go.sum ./

# Download dependencies and cache them
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the application port
EXPOSE 8080

# Run the Go application
CMD ["./main"]

