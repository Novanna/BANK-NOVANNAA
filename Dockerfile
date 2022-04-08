# Start from golang base image
FROM golang:alpine

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git 

# Setup folders
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Download all the dependencies
RUN go mod tidy -compat=1.17

# Build the Go app
RUN go build -o binary .

# Run the executable
ENTRYPOINT ["./binary"]
