# Start with a Go base image
FROM golang:1.24-alpine AS builder

# Set the working directory in the container
WORKDIR /app

# Copy the Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the entire application code to the container
COPY . .

# Build the Go application
RUN go build -o user-management-api ./app

# Start a new image for running the application (this keeps the final image smaller)
FROM alpine:latest

# Install dependencies for running the Go app
RUN apk --no-cache add ca-certificates

# Copy the built Go binary from the previous stage
COPY --from=builder /app/user-management-api /usr/local/bin/user-management-api

# Set environment variables (can also use .env file or Docker Compose for dynamic values)
ENV DB_HOST=postgres
ENV DB_PORT=5432
ENV DB_USER=user
ENV DB_PASSWORD=password
ENV DB_NAME=userdb

# Expose port 8080
EXPOSE 8080

# Command to run the Go application
CMD ["user-management-api"]
