# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory in the container
WORKDIR /go/src/app

# Copy the local package files to the container's workspace
COPY . .

# Build the application
RUN go mod tidy && go build -o app

# Expose the port on which the service will run
EXPOSE 8081

# Run the application
CMD ["./app"]
