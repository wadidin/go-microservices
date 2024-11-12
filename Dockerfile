# Use the official Go image
FROM golang:1.19-alpine

# Install gcc and other necessary build tools
RUN apk add --no-cache gcc musl-dev

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o hello-world-api

# Expose port 8070 to the outside world
EXPOSE 8070

# Command to run the executable
CMD ["./hello-world-api"]
