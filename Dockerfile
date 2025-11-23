FROM golang:1.25.3

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum first for better caching
COPY go.mod go.sum ./
RUN go mod tidy

# Copy all other files
COPY . .

# Build the Go binary
RUN go build -o hifi-proxy .

# Run the binary
CMD ["./hifi-proxy"]
