## To run the Postgres container, type:

```bash
docker compose up -d go_db
```

## To show all the containers (running and stopped ones) type:

```bash
docker ps -a
```

## Let's go back to the folder where the docker-compose.yml is located and type:

```bash
docker compose build
```

## Now, to check if the image has been built successfully, type:

```bash
docker images
```

## Run APP

```bash
docker compose up go-app
```