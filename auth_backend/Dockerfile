FROM golang:1.24.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o userAuth ./cmd/main.go

# Now, create a new lightweight image for running the application
FROM alpine:latest

WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/userAuth .

# Set environment variables
ENV PORT=8080

# Expose the port that the app will run on
EXPOSE 8080

# Run the application when the container starts
CMD ["./userAuth"]
