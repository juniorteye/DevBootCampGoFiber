# Use the official Golang image with Go 1.22.4
FROM golang:1.22.4 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .



# Use a minimal base image for the final image
FROM gcr.io/distroless/base-debian11


# Copy the compiled binary from the builder stage
COPY --from=builder /app/main /main
# Set the entrypoint to run the binary

# COPY /app/main / main
# CMD ["/main"]
ENTRYPOINT ["/main"]

