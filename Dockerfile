# Use the official Golang image as a build stage
FROM golang:1.23-alpine AS builder
ARG GROUP_ID=6085
ARG USER_ID=1654

RUN apk add --no-check-certificate git
RUN rm -rf /var/cache/apk/*

# Set the Current Working Directory inside the container
WORKDIR /app

RUN addgroup -g ${GROUP_ID} goappgroup && \
    adduser -D -u ${USER_ID} -G goappgroup goappuser

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

RUN chown -R goappuser:goappgroup /app
RUN chmod 777 -R /app
USER goappuser

# Download and install dependencies
RUN go mod tidy

# Copy the source code into the container
COPY . .

# go fmt
RUN go fmt -x ./...

# Build the Go app
RUN go build -buildvcs=false -o palko-htmx ./cmd

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./palko-htmx"]
