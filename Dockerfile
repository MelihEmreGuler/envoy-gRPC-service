# syntax=docker/dockerfile:1
FROM golang:1.21.1

# Set destination for COPY
WORKDIR /app/server

# Download Go modules
COPY go.mod go.sum  ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY ./ ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /server-built

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 50051

# Run
CMD ["/client-built"]

# Use a minimal base image to keep the image size small
FROM envoyproxy/envoy:v1.15.0 as envoy

# Copy your Envoy configuration file to the appropriate location in the container
COPY /envoy/envoy.yaml /envoy/envoy.yaml

# Expose the ports that Envoy will listen on (e.g., 8080 for HTTP)
EXPOSE 8080

# Specify the command to start Envoy with your configuration
CMD ["envoy", "-c", "/etc/envoy/envoy.yaml"]






