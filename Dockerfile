# Containerize the Go application that we have created
# This is the Dockerfile that we will use to build the image
# and run the container

# Start with a base image
FROM golang:1.21 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download all the dependencies
RUN go mod download

# Copy the source code to the working directory
COPY . .

# Build the application
RUN go build -o main .

#######################################################
# Reduce the image size using multi-stage builds
# We will use a distroless image to run the application
FROM gcr.io/distroless/base

# Copy the binary from the previous stage
COPY --from=builder /app/main .

# Copy the static files from the previous stage (if any)
COPY --from=builder /app/static ./static

# Expose the port on which the application will run
EXPOSE 80  # Change this to 80 since your application is intended to run on port 80

# Command to run the application
CMD ["./main"]
