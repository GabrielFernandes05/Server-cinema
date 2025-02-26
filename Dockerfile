# Base image
FROM golang:1.23-alpine

# Set working directory
WORKDIR /app

# Copy Go module files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o main .

# Expose the port and start the app
EXPOSE 8080
CMD ["./main"]