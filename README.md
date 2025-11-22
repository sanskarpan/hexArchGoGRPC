# Hexagonal Architecture Go gRPC Service

A production-ready microservice implementing basic arithmetic operations using hexagonal architecture, gRPC communication, and MySQL persistence.

## Table of Contents

- [Overview](#overview)
- [Architecture](#architecture)
- [Technology Stack](#technology-stack)
- [Project Structure](#project-structure)
- [API Reference](#api-reference)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [Testing](#testing)
- [Docker Deployment](#docker-deployment)
- [Development](#development)
- [Contributing](#contributing)

## Overview

This project demonstrates a clean implementation of **hexagonal architecture** (also known as ports and adapters architecture) in Go. It provides a gRPC service for basic arithmetic operations (addition, subtraction, multiplication, division) with persistent storage of operation history in MySQL.

### Key Features

- **Hexagonal Architecture**: Clear separation of concerns with core business logic isolated from external dependencies
- **gRPC Communication**: High-performance RPC communication using Protocol Buffers
- **Persistent Storage**: Operation history stored in MySQL database
- **Comprehensive Testing**: Unit tests for core logic and integration tests for gRPC endpoints
- **Docker Support**: Containerized deployment with docker-compose
- **Dependency Injection**: Proper dependency injection following hexagonal architecture principles

## Architecture

The project follows hexagonal architecture principles:

```
┌─────────────────────────────────────────────────────────────────┐
│                    External Clients (gRPC)                      │
└─────────────────────┬───────────────────────────────────────────┘
                      │ Driving Adapters
                      ▼
┌─────────────────────────────────────────────────────────────────┐
│                   Application Layer (API)                       │
│  - Orchestrates operations between core and adapters               │
│  - Implements use cases                                         │
│  - Dependency injection                                         │
└─────────────────────┬───────────────────────────────────────────┘
                      │ Ports
                      ▼
┌─────────────────────────────────────────────────────────────────┐
│                     Core Business Logic                         │
│  - Pure business logic (no external dependencies)               │
│  - Domain models and business rules                             │
└─────────────────────┬───────────────────────────────────────────┘
                      │ Ports
                      ▼
┌─────────────────────────────────────────────────────────────────┐
│                  Driven Adapters (Database)                     │
│  - Database operations                                          │
│  - External service integrations                                │
└─────────────────────────────────────────────────────────────────┘
```

### Architecture Benefits

- **Testability**: Core business logic can be tested in isolation
- **Flexibility**: Easy to swap implementations (e.g., different databases)
- **Maintainability**: Clear separation of concerns
- **Technology Agnostic**: Core logic independent of external frameworks

## Technology Stack

- **Language**: Go 1.15+
- **Communication**: gRPC with Protocol Buffers
- **Database**: MySQL 8.0+
- **Query Builder**: Squirrel SQL builder
- **Testing**: Testify framework
- **Containerization**: Docker & Docker Compose
- **Architecture**: Hexagonal Architecture

## Project Structure

```
hexArchGoGRPC/
├── cmd/
│   └── main.go                    # Application entry point
├── internal/
│   ├── adapters/
│   │   ├── framework/
│   │   │   ├── left/grpc/         # gRPC adapter (driving adapter)
│   │   │   │   ├── pb/           # Generated protobuf code
│   │   │   │   ├── proto/         # Protocol buffer definitions
│   │   │   │   ├── rpc.go         # gRPC service implementation
│   │   │   │   ├── server.go      # gRPC server setup
│   │   │   │   └── rpc_test.go    # gRPC integration tests
│   │   │   └── right/db/          # Database adapter (driven adapter)
│   │   │       └── db.go          # Database operations
│   ├── application/
│   │   ├── api/
│   │   │   ├── api.go             # Application layer implementation
│   │   │   └── interfaces.go      # Core business logic interface
│   │   └── core/arithmetic/
│   │       ├── arithmetic.go      # Core arithmetic operations
│   │       └── arithmetic_test.go # Unit tests
│   └── ports/
│       ├── left.go                # Driving port (API interface)
│       └── right.go               # Driven port (Database interface)
├── testdb/
│   └── init.sql                   # Database schema initialization
├── docker-compose.yaml            # Multi-container setup
├── Dockerfile                     # Application containerization
├── grpc_entrypoint.sh             # Container startup script
├── go.mod                         # Go module dependencies
├── go.sum                         # Dependency checksums
└── README.md                      # Project documentation
```

## API Reference

### gRPC Service: ArithmeticService

The service provides four arithmetic operations:

#### GetAddition
Adds two numbers and stores the result in history.

**Request:**
```protobuf
message OperationParameters {
  int32 a = 1;
  int32 b = 2;
}
```

**Response:**
```protobuf
message Answer {
  int32 value = 1;
}
```

#### GetSubtraction
Subtracts second number from first and stores the result in history.

#### GetMultiplication
Multiplies two numbers and stores the result in history.

#### GetDivision
Divides first number by second and stores the result in history.

### Example Usage

```bash
# Using grpcurl
grpcurl -plaintext -d '{"a": 10, "b": 5}' localhost:9000 pb.ArithmeticService.GetAddition
```

## Prerequisites

- Go 1.15 or higher
- MySQL 8.0 or higher
- Docker and Docker Compose (optional, for containerized deployment)
- Protocol Buffer compiler (`protoc`)

## Installation

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd hexArchGoGRPC
   ```

2. **Install dependencies:**
   ```bash
   go mod download
   ```

3. **Generate Protocol Buffer code:**
   ```bash
   # Install protoc if not already installed
   # Ubuntu/Debian:
   sudo apt-get install protobuf-compiler

   # Generate Go code from .proto files
   cd internal/adapters/framework/left/grpc/proto
   protoc --go_out=. --go-grpc_out=. *.proto
   ```

4. **Set up the database:**
   ```bash
   # Start MySQL server
   # Create database 'hex_test' with user 'root' and password 'Admin123'
   # The init.sql script will create the required tables
   ```

## Configuration

Configure the following environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| `DB_DRIVER` | Database driver | `mysql` |
| `DS_NAME` | Database connection string | `root:Admin123@tcp(localhost:3306)/hex_test` |
| `MYSQL_HOST` | MySQL host | `localhost` |
| `MYSQL_PORT` | MySQL port | `3306` |
| `MYSQL_USER` | MySQL username | `root` |
| `MYSQL_PASSWORD` | MySQL password | `Admin123` |
| `MYSQL_DB` | MySQL database name | `hex_test` |

## Usage

### Running Locally

1. **Start MySQL server** (if not using Docker)

2. **Set environment variables:**
   ```bash
   export DB_DRIVER=mysql
   export DS_NAME="root:Admin123@tcp(localhost:3306)/hex_test"
   ```

3. **Run the application:**
   ```bash
   go run cmd/main.go
   ```

4. **The gRPC server will start on port 9000**

### Testing the Service

```bash
# Test core arithmetic operations
go test ./internal/application/core/arithmetic/

# Test gRPC endpoints (requires running server)
go test ./internal/adapters/framework/left/grpc/

# Run all tests
go test -v ./...
```

## Docker Deployment

### Using Docker Compose (Recommended)

1. **Start the complete stack:**
   ```bash
   docker-compose up -d
   ```

2. **This will start:**
   - MySQL database on port 3307 (mapped from 3306)
   - gRPC service with automatic database initialization

3. **View logs:**
   ```bash
   docker-compose logs -f grpc
   ```

4. **Stop the stack:**
   ```bash
   docker-compose down
   ```

### Manual Docker Build

1. **Build the image:**
   ```bash
   docker build -t hex-grpc-service .
   ```

2. **Run the container:**
   ```bash
   docker run -d \
     --name hex-grpc \
     -e DB_DRIVER=mysql \
     -e DS_NAME="root:Admin123@tcp(host.docker.internal:3306)/hex_test" \
     -p 9000:9000 \
     hex-grpc-service
   ```

## Testing

### Unit Tests

```bash
# Test core arithmetic logic
go test ./internal/application/core/arithmetic/

# Test with coverage
go test -cover ./internal/application/core/arithmetic/
```

### Integration Tests

```bash
# Test gRPC endpoints (requires running database)
go test ./internal/adapters/framework/left/grpc/
```

### Manual Testing with grpcurl

1. **Install grpcurl:**
   ```bash
   go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
   ```

2. **Test addition:**
   ```bash
   grpcurl -plaintext -d '{"a": 10, "b": 5}' localhost:9000 pb.ArithmeticService.GetAddition
   ```

3. **Test other operations:**
   ```bash
   grpcurl -plaintext -d '{"a": 10, "b": 5}' localhost:9000 pb.ArithmeticService.GetSubtraction
   grpcurl -plaintext -d '{"a": 10, "b": 5}' localhost:9000 pb.ArithmeticService.GetMultiplication
   grpcurl -plaintext -d '{"a": 10, "b": 5}' localhost:9000 pb.ArithmeticService.GetDivision
   ```

## Development

### Adding New Operations

1. **Add to core business logic:**
   ```go
   // In internal/application/core/arithmetic/arithmetic.go
   func (arith Arith) NewOperation(a, b int32) (int32, error) {
       return a + b, nil // Your logic here
   }
   ```

2. **Add to application interface:**
   ```go
   // In internal/application/api/interfaces.go
   type Arithmetic interface {
       NewOperation(a, b int32) (int32, error)
   }
   ```

3. **Implement in application layer:**
   ```go
   // In internal/application/api/api.go
   func (apia Application) GetNewOperation(a, b int32) (int32, error) {
       answer, err := apia.arith.NewOperation(a, b)
       if err != nil {
           return 0, err
       }
       err = apia.db.AddToHistory(answer, "new_operation")
       return answer, err
   }
   ```

4. **Add to gRPC service:**
   ```protobuf
   // In proto/arithmetic_svc.proto
   service ArithmeticService {
       // ... existing operations
       rpc GetNewOperation(OperationParameters) returns (Answer) {}
   }
   ```

5. **Implement gRPC handler:**
   ```go
   // In internal/adapters/framework/left/grpc/rpc.go
   func (grpca Adapter) GetNewOperation(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
       // Implementation similar to existing operations
   }
   ```

### Database Schema Changes

Update `testdb/init.sql` and recreate the database:

```bash
docker-compose down -v  # Remove volumes
docker-compose up -d     # Recreate with new schema
```

## Acknowledgments

- **Alistair Cockburn** for defining hexagonal architecture
- **Go gRPC team** for the excellent gRPC implementation
- **Squirrel SQL** for the query builder library
