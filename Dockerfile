# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy all files from the current directory (.) into the working directory (/app) inside the container
COPY . .

# Copy the database file into the container
COPY ./database/database.db ./cmd/database/database.db
COPY ./uploads/photo ./cmd/uploads/photo
# Build your application (assuming main.go is located inside the cmd directory)
RUN go build -o main ./cmd

# Expose the port that your application will use
EXPOSE 8082

# Command to run when the container starts
CMD ["./main"]

