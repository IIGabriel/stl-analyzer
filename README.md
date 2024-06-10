# STL File Analysis API

This API allows users to analyze STL files. It is built using Go and provides endpoints for interacting with and analyzing STL files. Runs on port 8080.

## Getting Started

### Prerequisites

- Go 1.22

### Running the Project

To run the project, use the following command:

```bash
go run cmd/api/main.go
```

The API will be available at `http://localhost:8080`.

### Swagger Documentation

Swagger is used to provide interactive API documentation. To initialize Swagger, use the following command:

Para instalar o Swagger, utilize o seguinte comando:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

```bash
swag init --parseDependency --parseInternal
```

Once Swagger is initialized, you can access the documentation at `http://localhost:8080/swagger/*`.

### Mock Generation

Mocks are used for testing purposes. To generate mocks, use the following command:

```bash
go install github.com/golang/mock/mockgen@latest
```
```bash
mockgen -source=internal/stl/usecase.go -destination=internal/stl/mocks/usecase_mock.go -package=mock
```

### Running All Tests

To run all tests in the project, use the following command:

```bash
go test ./...
```

## API Endpoints

- `/stl/triangle` - Endpoint to analyze STL files.
- `/swagger/*` - Swagger documentation.
