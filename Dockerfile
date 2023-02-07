# Use the latest Alpine image as the build stage
FROM golang:alpine AS builder

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Build the Go binary
RUN go build -o disk-space-monitor .

# Use the smallest Alpine image as the final stage
FROM alpine

# Set the working directory to /app
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/disk-space-monitor .

# Run the binary as a daemon
CMD ["./disk-space-monitor", "--interval", "30", "--path", "/", "--threshold", "90"]
