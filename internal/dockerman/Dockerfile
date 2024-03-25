# Use the official Go image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the contents of ../../advocate-2 to /app in the container
COPY ../../advocate-2 .

CMD ["go", "run", "main.go"]