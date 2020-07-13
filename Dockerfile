# Get golang's alpines image as the base image
FROM golang:alpine

# Set working directory inside container
WORKDIR /usr/src/app

# Copy .mod and .sum files to the working directory inside container
COPY go.mod go.sum ./

# Run go mod to download all dependancies
RUN go mod download

# Copy entire soucre code from host to the working directory inside container
COPY  ./ ./

# Get compiledaemon to rebuild and restart the application
RUN go get github.com/githubnemo/CompileDaemon

# Expose port 3000 to the host machine
EXPOSE 3000

# Configure compiledaemon to rebuild and restart the application
CMD CompileDaemon --build="go build omelette.go" --command=./omelette



