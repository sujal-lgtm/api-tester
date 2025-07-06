# Use minimal Go image
FROM golang:1.21-alpine

# Set working directory
WORKDIR /app

# Copy go.mod and download dependencies
COPY go.mod ./
RUN go mod download

# Copy the entire source code
COPY . .

# Build the main file inside cmd/
RUN go build -o server ./cmd/main.go

# Expose the port (used by Render)
EXPOSE 6969

# Start the server
CMD ["./server"]
