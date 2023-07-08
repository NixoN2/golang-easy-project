# Start from the official Go image
FROM golang:1.20-alpine

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download and install project dependencies
RUN go mod download

# Copy the rest of the project files
COPY . .

# Build the Go application
RUN go build -o main ./cmd/app/main.go

# Set the command to run the executable
CMD ["./main"]