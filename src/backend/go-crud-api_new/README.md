# Go-CRUD-API

## Project Overview

This project is a Go-based CRUD API with services that use PostgreSQL, Redis, and InfluxDB for data storage and
management.

## Prerequisites

Before running the project, make sure you have the following installed:

- [Go](https://golang.org/dl/) (Version 1.16 or higher)
- [Docker](https://www.docker.com/get-started) (to run PostgreSQL, Redis, and InfluxDB containers)
- [Git](https://git-scm.com/) (optional but recommended)

### Clone the Repository

```bash
git clone https://github.com/your-repo/go-crud-api.git
cd go-crud-api
```

## Environment Variables

Ensure you have the necessary environment variables set for your application. You can use an `.env` file to store them.

```plaintext
DATABASE_URL=postgres://admin:admin_password@localhost:5432/app_db
REDIS_URL=redis://localhost:6379/0
INFLUXDB_URL=http://localhost:8086
INFLUXDB_TOKEN=your_token
```

## Starting the Application

### 1. Docker Setup

You need to make sure Docker is running, as the application relies on Docker containers for PostgreSQL, Redis, and
InfluxDB.

Use the following command to start the required containers:

```bash
docker-compose up -d
```

This command will start the following containers:

- **app_redis** (Redis)
- **app_influxdb** (InfluxDB)
- **app_postgres** (PostgreSQL)

### 2. Running the Application

Instead of running an `.exe` file (e.g., `main.exe`), use `go run` to execute your Go application directly from source.

To run the application, use:

```bash
go run main.go
```

This will start the application and ensure that all the services (PostgreSQL, Redis, InfluxDB) are initialized.

### 3. Rebuild Docker Images

If you need to rebuild the Docker images from scratch, use:

```bash
docker-compose build --no-cache
```

This ensures that Docker does not use cached layers and rebuilds all images.

### 4. View Running Docker Containers

To verify that Docker containers are running, use:

```bash
docker ps
```

This lists all running containers and their status.

### 5. Accessing Redis and InfluxDB

#### Redis:

- **CLI Access**: You can interact with Redis using the Redis CLI:

  ```bash
  docker exec -it app_redis redis-cli
  ```

#### InfluxDB:

- **Web Interface**: Access the InfluxDB web interface at:

  ```plaintext
  http://localhost:8086
  ```

Use the interface to manage databases and monitor data.

### 6. Running Tests

#### Run Unit Tests

To run the unit tests, use:

```bash
go test -v ./tests
```

#### Check Test Coverage

To check the test coverage, use:

```bash
go test -cover ./tests
```

#### Generate a Coverage Report

To generate a detailed coverage report:

```bash
go test -coverprofile=coverage.out ./tests
go tool cover -html=coverage.out
```

This opens an HTML report in your browser showing which lines of code are covered by tests.

---

By following the steps above, you can start and manage your Go-based application without the need to manually execute
`main.exe`. Instead, use `go run main.go` for easier and direct execution.