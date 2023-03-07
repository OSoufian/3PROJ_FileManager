# Use the official Go image as the base image
FROM golang:1.17-alpine AS build

# Set the working directory to /app
WORKDIR /app

# Copy the source code to the container
COPY . .

# Build the binary
RUN go build -o main .

# Use a lightweight image for the final stage
FROM alpine:latest

# Copy the binary from the previous stage
COPY --from=build /app/main .

# Set the working directory to /app
WORKDIR /app

# Install necessary packages
RUN apk add --no-cache ca-certificates

# Expose the port that the application will listen on
EXPOSE 3000

# Run the binary
CMD ["./main"]
