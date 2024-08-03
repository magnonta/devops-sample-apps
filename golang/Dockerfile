# Stage 1: Build the Go binary
FROM golang:1.20-alpine AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go source code
COPY . .

# Set the necessary environment variables for cross-compilation
ENV GOARCH=amd64
ENV GOOS=linux

# Build the Go app for the correct architecture
RUN go build -o main .

# Stage 2: A minimal Docker image to run the Go binary
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the pre-built binary file from the previous stage
COPY --from=build /app/* .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
