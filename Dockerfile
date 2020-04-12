# Get golang's alpines image as your base image
FROM golang:alpine

# Set working directory inside container
WORKDIR /usr/src/app

# Copy .mod and .sum files to the working directory inside container
COPY go.mod go.sum ./

# Run go mod to download all dependancies
RUN go mod download

# Copy entire soucre code from host to the working directory inside container
COPY  ./ ./

# Get compiledaemon to rebuild and restart my application
RUN go get github.com/githubnemo/CompileDaemon

# Expose port 8080 to the outside world
EXPOSE 3000

# Configure compiledaemon to rebuild and restart my application
CMD CompileDaemon --build="go build server.go" --command=./server



