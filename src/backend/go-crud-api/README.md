## Instructions for Running the Project

### 1. Start the PostgreSQL Container

To start the PostgreSQL container, run the following command:

```bash
docker compose up -d go_db
```

This will start the PostgreSQL container in detached mode.

### 2. View All Containers

To display all Docker containers (both running and stopped), use the following command:

```bash
docker ps -a
```

### 3. Build the Docker Image

Navigate to the directory where your `docker-compose.yml` file is located, and then build the Docker image with the following command:

```bash
docker compose build
```

### 4. Verify the Docker Image

After building, check if the image has been successfully created by running:

```bash
docker images
```

### 5. Run the Application

To start the application, execute the following command:

```bash
docker compose up go-app
```

This will launch the application container.

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
