# Adapptor Backend API

Technologies: Golang, Docker

I decided to build the backend server with standard Go libs as well as gorilla mux for the routing.
Using swaggo for OpenAPI documentation. I also used an assertion library for testing.

## Features

The backend has a simple HTTP server listening on localhost:8080 for incoming requests.

I have added two simple middlewares:

- logging middleware for incoming requests
- auth middleware for simple Bearer token authentication, can be applied granularly at an endpoint level

I chose this directory structure as it keeps the `main.go` fairly lean, and separates the server/middleware/endpoint code for easier unit testing.

## Improvements/Choices

A point of improvement would be using the latest version (v3) of OpenAPI schema generation, as well as code generation based on that spec. I decided to use v2 since it was more straightfoward to setup as well as progress with.

Another point of improvement would be a more stricter request/response pattern for all endpoints - ideally governed via a middleware. So the endpoints would need to return a response or an error only, and the responsibility of writing the response back in a standard structure would fall on a middleware.

Also an easy win for Go HTTP server is a panic handler middleware. This prevents the entire server from crashing from a failed request.

## Prerequisites

- Go 1.23
- Docker

## Running the Application with Docker

1. Install dependencies:

```bash
go mod download
```

2. Generate Swagger documentation:

```bash
# Install swag if you haven't already
go install github.com/swaggo/swag/cmd/swag@latest

# Generate docs
swag init -g ./cmd/main.go
```

3. Build the docker container:

```bash
docker build -t adapptor-backend .
```

4. Run the docker container:

```bash
docker run -p 8080:8080 adapptor-backend
```

## API Documentation

Once the server is running, you can access the Swagger UI documentation at:

```
http://localhost:8080/swagger/
```

## Testing

There are unit tests for the endpoint implementations as well as an integration test to test the API as a client making a request.

Run the tests with the following command at the root of the repo:

```bash
go test -v ./...
```
