# Use the official Golang image as a base
FROM golang:1.22.2 AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application with static linking
RUN CGO_ENABLED=0 GOOS=linux go build -o fleet-api -a -installsuffix cgo cmd/fleet-api/main.go

# Start a new stage from scratch
FROM scratch

# Set the current working directory inside the container
WORKDIR /app

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/fleet-api .

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./fleet-api"]
