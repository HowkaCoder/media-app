# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app
RUN apt-get update && \
    apt-get install -y libwebp-dev && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Copy all files from the current directory (.) into the working directory (/app) inside the container
COPY . .


# Build your application (assuming main.go is located inside the cmd directory)
RUN go build -o main . 

# Expose the port that your application will use
EXPOSE 8082

# Command to run when the container starts
CMD ["./main"]

