# Access System API

A backend API for access control, written in Go.

## Status

MVP

## Features

- Layered Architecture
- Dockerized for easy deployment
- Database initialization scripts
- Unit and integration tests

## Prerequisites

- Go 1.20+
- Docker & Docker Compose
- Mockgen

## Getting Started

### Clone the repository

```
git clone https://github.com/iwanttobelievee/access-system-api.git
cd access-system-api
```

### Build and run with Docker Compose

```
go mod download
sudo chmod +x run.sh
./run.sh dev
```

The API will be available at `http://localhost:8081/`.
PgAdmin will be available at `http://localhost:8091/`.

### 
### Running Tests

To run all tests:

```
./run.sh test
```

## Endpoints

### POST /api/v1/embedding
Add embedding.
Parameters:
- `name` (string, required): Name of the embedding.
- `vector` (array of float32, required, 512): Embedding vector.

### POST /api/v1/embedding/validate
Validate embedding.
Parameters:
- `vector` (array of float32, required, 512): Embedding vector.

### DELETE /api/v1/embedding
Delete embedding.
Parameters:
- `id` (int64, required): ID of the embedding to delete.

## Project Structure

- `cmd/` - Entry point (main.go)
- `internal/` - Application logic
  - `cfg/` - Configuration
  - `client/` - External clients
  - `domain/` - Domain models
  - `handler/` - HTTP handlers
  - `mocks/` - Test mocks
  - `repository/` - Data access
  - `router/` - Routing
  - `service/` - Business logic
- `docker/` - Dockerfiles and DB scripts

## Database

Database initialization scripts are in `docker/db/scripts/init.sql`.

## Environment Configuration

Create a `.env` file in the project root to configure database and admin credentials.

You can copy `.env` from `.env.example` and adjust values as needed for your environment.

## License

MIT
