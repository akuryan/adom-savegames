# Use the official Go image as the base image
ARG builder=golang:1.20
FROM ${builder} AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# execute tests
RUN go test ./... -v

# Build the Go application
RUN go build -o adom-savegames ./cmd