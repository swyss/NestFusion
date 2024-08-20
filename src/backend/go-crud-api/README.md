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

### 5.5. Reset Docker Containers and Images

Before rebuilding your Docker containers and images after making changes to your application, you should reset your
Docker environment by following these steps:

#### 5.5.1 Stop and Remove All Running Containers

To stop and remove all running containers, run the following command:

```bash
docker stop $(docker ps -q) && docker rm $(docker ps -aq)
```

This command stops all running containers and removes all containers, both running and stopped.

#### 5.5.2 Remove Old Docker Images

To ensure that old images that might still be in use are removed, run the following command:

```bash
docker rmi $(docker images -q)
```

This command removes all Docker images from your system. Make sure you don't need any images that should not be rebuilt
before running this command.

#### 5.5.3 Remove Unused Docker Networks and Volumes

To ensure that old networks and volumes that are no longer needed are removed, run the following commands:

```bash
docker network prune -f
docker volume prune -f
```

These commands remove all unused networks and volumes to free up space and avoid potential conflicts.

#### 5.5.4 Remove All Old Docker Compositions

If you want to ensure that old Docker Compose configurations are removed before restarting, you can run the following
command:

```bash
docker compose down --rmi all --volumes --remove-orphans
```

This command removes all containers, images, volumes, and networks created by `docker-compose`.

### 6. Rebuild the Docker Images and Start the Containers

After resetting the old instances, you can rebuild the Docker images and start the containers with the following
commands:

```bash
docker compose build --no-cache
docker compose up -d
```

The `--no-cache` option ensures that no cached layers are used, ensuring that all dependencies and code are built from
scratch.

### 7. Verify the Docker Image

After building, check if the image has been successfully created by running:

```bash
docker images
```

### 8. Run the Application

To start the application, execute the following command:

```bash
# Stop
docker stop $(docker ps -q)
# Run
docker compose up --build
docker compose up go_app
```

This will launch the application container.

### 9. Generate Swagger Documentation

To generate the Swagger documentation, run:

```bash
swag init
```

After generating the documentation, you can view it in your browser at:

```plaintext
http://localhost:8000/swagger/index.html
```

This URL will display the Swagger UI, where you can interact with the API documentation.

### 10. Accessing Redis and InfluxDB

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

### 11. Running Tests

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

Here’s the translated and detailed explanation for the commands in English:

---

### Reset Docker Containers and Start the Application

To ensure that your Docker environment is reset and your application runs with the latest changes, follow these steps:

```bash
# 1. Stop all running Docker containers
docker stop $(docker ps -q)
```

This command stops all currently running Docker containers. The `docker ps -q` command lists the IDs of all running
containers, and `docker stop` stops them.

```bash
# 2. Remove all Docker containers
docker rm $(docker ps -aq)
```

This command removes all Docker containers from your system. The `docker ps -aq` command lists the IDs of all
containers (both running and stopped), and `docker rm` removes them.

```bash
# 3. Remove all Docker images
docker rmi $(docker images -q)
```

This command removes all Docker images from your system. The `docker images -q` command lists the IDs of all images, and
`docker rmi` removes them.

```bash
# 4. Remove all unused Docker networks
docker network prune -f
```

This command removes all Docker networks that are not currently in use. The `-f` flag forces the removal without
prompting for confirmation.

```bash
# 5. Remove all unused Docker volumes
docker volume prune -f
```

This command removes all Docker volumes that are not currently in use. The `-f` flag forces the removal without
prompting for confirmation.

```bash
# 6. Remove all Docker-Compose configurations (containers, images, volumes, networks)
docker compose down --rmi all --volumes --remove-orphans
```

This command stops and removes all containers, images, volumes, and networks created by Docker Compose. The `--rmi all`
flag removes all images used by any service, `--volumes` removes named volumes declared in the `volumes` section of the
Compose file, and `--remove-orphans` removes containers not defined in the Compose file but still attached to the same
network.

```bash
# 7. Rebuild Docker images without using the cache
docker compose build --no-cache
```

This command rebuilds the Docker images, ensuring that no cached layers are used. This forces Docker to build everything
from scratch, which can be important if there have been changes to the Dockerfile or dependencies.

```bash
# 8. Start the Docker containers in the background
docker compose up -d
```

This command starts the Docker containers defined in your `docker-compose.yml` file in detached mode, meaning they will
run in the background.

```bash
# 9. Verify that all containers are running
docker ps
```

This command lists all running Docker containers. It helps verify that all containers started correctly after the reset.

### Start the Application

After resetting the Docker environment and ensuring that everything is clean, you can start your application with the
following commands:

```bash
# Stop any currently running containers
docker stop $(docker ps -q)

# Rebuild and start the application
docker compose up --build
docker compose up go_app
```

- The first `docker stop` command ensures that any previously running containers are stopped.
- `docker compose up --build` rebuilds the Docker images and starts the containers with the latest changes.
- `docker compose up go_app` specifically starts the application container.

These steps will help you reset your Docker environment and ensure that your application is running with the latest
updates.

```bash
# 1. Stoppe alle laufenden Docker-Container
docker stop $(docker ps -q)

# 2. Entferne alle Docker-Container
docker rm $(docker ps -aq)

# 3. Entferne alle Docker-Images
docker rmi $(docker images -q)

# 4. Entferne alle nicht verwendeten Docker-Netzwerke
docker network prune -f

# 5. Entferne alle nicht verwendeten Docker-Volumes
docker volume prune -f

# 6. Entferne alle Docker-Compose-Konfigurationen (Container, Images, Volumes, Netzwerke)
docker compose down --rmi all --volumes --remove-orphans

# 7. Baue die Docker-Images neu ohne Cache
docker compose build --no-cache

# 8. Starte die Docker-Container neu im Hintergrund
docker compose up -d

# 9. Verifiziere, dass alle Container laufen
docker ps


# Stop any currently running containers
docker stop $(docker ps -q)

# Rebuild and start the application
docker compose up --build
docker compose up go_app
```

```bash
# Stop any currently running containers
docker stop $(docker ps -q)

# Rebuild and start the application
docker compose up --build
docker compose up go_app
```
