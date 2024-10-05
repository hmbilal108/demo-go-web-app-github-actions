# Use a base image for building the Go application
FROM golang:1.21 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the source code to the working directory
COPY . .

# Build the application
RUN go build -o main .

# Use a distroless image to reduce the size of the final image
FROM gcr.io/distroless/base

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Copy the static files (if any)
COPY --from=builder /app/static ./static

# Expose the port on which the application will run
EXPOSE 80

# Command to run the application
CMD ["./main"]
