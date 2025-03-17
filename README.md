# Modular Monolith Sample

This repository demonstrates a modular monolith architecture using Go. It is designed to showcase how large-scale applications can be structured into well-defined modules.

## Overview

- **Modular Design:** Organize complex applications into separate modules.
- **Clean Architecture:** Maintain a clear separation between domain, application, and infrastructure layers.
- **Technology Stack:**
  - Go (Golang)
  - MySQL as the database
  - SQLX for database interactions
- **Inter-module Communication:**
  - Modules communicate using gRPC.
  - The main interface utilizes a REST API built with GIN.

## Modules

- **User Module:**
  Contains functionalities related to user management, along with its own infrastructure and domain logic.

<!-- ...add additional module descriptions as needed... -->

## Getting Started

1. Clone the repository.
2. Install the required dependencies.
3. **Generate Protocol Buffers:**
   - Use buf to generate the proto files:
     ```bash
     cd proto
     buf generate
     ```
4. **Run Docker Compose:**
   - Start the necessary services:
     ```bash
     docker-compose up
     ```
5. **Run the Application:**
   - Start the API server:
     ```bash
     go run cmd/api/main.go
     ```

For more details on setup and usage, please refer to the corresponding module documentation.

## Migration

The migration tool uses [golang-migrate/migrate](https://github.com/golang-migrate/migrate)

```bash
docker run -v $PWD/schema/migrations{{ target }}:/migrations migrate/migrate create -ext sql -dir /migrations -seq foobar
```

#### Table Creation

```bash
docker run -v $PWD/schema/migrations/{{ target }}:/migrations --network host migrate/migrate \
  -path=/migrations -database "mysql://root@tcp(localhost:3306)/{{ target }}" up
```
