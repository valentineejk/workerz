FROM golang:1.22.2-alpine3.19 as builder


RUN mkdir /app
ADD . /app
WORKDIR /app


# Build the Go application
RUN CGO_ENABLED=0 go build -o workerz .
RUN chmod +x /app/workerz

# Final lightweight image
FROM alpine:latest

# Create a directory for the application
RUN mkdir /app

# Copy the executable from the builder stage to the final image
COPY --from=builder /app/workerz /app

# Copy the .env file into the final image
COPY .env /app/.env

# Set the working directory
WORKDIR /app

# Command to run the executable
CMD ["./workerz"]


