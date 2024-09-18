## Instructions for Running the Project

### 1. Start the Application with Docker

There are different workflows depending on whether you need to reset Docker or not.

#### Workflow 1: Reset Docker Environment and Start Application

To reset your Docker environment and start the application, run the following:

```bash
go run main.go --reset-docker
```

This will stop all running containers, remove old images and volumes, reset the Docker environment, and start the
application.

#### Workflow 2: Start Application Without Docker Reset

If Docker containers are already running, and you only need to start the application, use:

```bash
go run main.go --app-only
```

This will skip the Docker setup and directly start the application.

#### Workflow 3: Check Docker and Start Application

If you want to check whether Docker containers are running before starting the application, simply run:

```bash
go run main.go
```

This will check if the necessary containers (PostgreSQL, Redis, InfluxDB) are running and start them if required. Then,
it starts the application.

### 2. Rebuild Docker Images

If you need to rebuild the Docker images from scratch, you can run the following command:

```bash
docker compose build --no-cache
```

This ensures that Docker does not use cached layers and rebuilds all images.

### 3. View Running Docker Containers

To verify that Docker containers are running, use the following command:

```bash
docker ps
```

This lists all running containers and their status.

### 4. Accessing Redis and InfluxDB

#### Redis:

- **CLI Access**: You can interact with Redis using the Redis CLI:

  ```bash
  docker exec -it go_redis redis-cli
  ```

#### InfluxDB:

- **Web Interface**: Access the InfluxDB web interface at:

  ```plaintext
  http://localhost:8086
  ```

Use the interface to manage databases and monitor data.

### 5. Running Tests

#### Run Unit Tests

To run the unit tests, use the following command:

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

### 6. Generate Swagger Documentation

To generate the Swagger documentation, run:

```bash
swag init
```

After generating the documentation, you can view it in your browser at:

```plaintext
http://localhost:8000/swagger/index.html
```

This URL will display the Swagger UI, where you can interact with the API documentation.

---

By following the steps above, you can reset, start, or manage your application and its Docker containers effectively,
ensuring everything runs smoothly.

```