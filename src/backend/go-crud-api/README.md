## Instructions for Running the Project

### 1. Start the PostgreSQL Container

To start the PostgreSQL container, run the following command:

```bash
docker compose up -d go_db
```

This will start the PostgreSQL container in detached mode.

### 2. Start the Redis Container

To start the Redis container, run the following command:

```bash
docker compose up -d go_redis
```

Redis will start in detached mode, and you can access the Redis CLI by running:

```bash
docker exec -it go_redis redis-cli
```

### 3. Start the InfluxDB Container

To start the InfluxDB container, run the following command:

```bash
docker compose up -d go_influxdb
```

Once started, you can access the InfluxDB web interface by navigating to:

```plaintext
http://localhost:8086
```

Log in with the credentials you set during the InfluxDB setup or the default credentials if this is a new setup.

### 4. View All Containers

To display all Docker containers (both running and stopped), use the following command:

```bash
docker ps -a
```

### 5. Build the Docker Image

Navigate to the directory where your `docker-compose.yml` file is located, and then build the Docker image with the
following command:

```bash
docker compose build
```

### 6. Verify the Docker Image

After building, check if the image has been successfully created by running:

```bash
docker images
```

### 7. Run the Application

To start the application, execute the following command:

```bash
docker compose up --build
docker compose up go-app
```

This will launch the application container.

### 8. Generate Swagger Documentation

To generate the Swagger documentation, run:

```bash
swag init
```

After generating the documentation, you can view it in your browser at:

```plaintext
http://localhost:8000/swagger/index.html
```

This URL will display the Swagger UI, where you can interact with the API documentation.

### 9. Accessing Redis and InfluxDB

#### Redis:

- **CLI Access**: You can interact with Redis using the Redis CLI:

  ```bash
  docker exec -it go_redis redis-cli
  ```

    - Use commands like `PING`, `SET`, `GET`, etc., to interact with the Redis store.

#### InfluxDB:

- **Web Interface**: Access the InfluxDB web interface at:

  ```plaintext
  http://localhost:8086
  ```

    - Use this interface to manage databases, view dashboards, and monitor data.

### 10. Running Tests

#### Run Unit Tests

To run the unit tests, use the following command:

```bash
go test -v ./tests
```

This will execute all the tests in the project and provide detailed output about which tests passed and which failed.

#### Check Test Coverage

To check the test coverage of your code, use:

```bash
go test -cover ./tests
```

This command will show the percentage of your code that is covered by tests.

#### Generate a Coverage Report

To generate a detailed coverage report, use the following commands:

```bash
go test -coverprofile=coverage.out ./tests
go tool cover -html=coverage.out
```

This will open an HTML report in your browser, showing which lines of code are covered by tests.