# A simple Todo application built with Go and Gin

Model, View, Service (MVS) architecture.

Separate the router file and call it from `main.go`.

## Usage

```sh
# POST /todos
$ curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"title":"勉強する","completed":false}'

# GET /todos
$ curl http://localhost:8080/todos

# DELETE /todos/:id
$ curl -X DELETE http://localhost:8080/todos/1

# PUT /todos/:id/toggle
$ curl -X PUT http://localhost:8080/todos/1/toggle

# GET /todos/completed
$ curl http://localhost:8080/todos/completed
```

## Testing

```sh
# Run all tests
$ go test ./...
# Run tests for a specific package
$ go test ./services
# Run tests with coverage
$ go test -cover ./...
# Run tests with verbose output
$ go test -v ./...
# Run tests
$ go test -run TestCreateAndGetTodos ./services
```
