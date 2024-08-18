package startup

import (
	"database/sql"
	"github.com/go-redis/redis/v8"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"go-crud-api/internal/databases"
	"go-crud-api/internal/databases/influxdb"
	"go-crud-api/internal/databases/redisdb"
	"go-crud-api/internal/logger"
)

// InitializeServices initializes all services and returns their instances.
func InitializeServices(l *logger.Logger) (*sql.DB, *redis.Client, influxdb2.Client) {
	l.InfoMsg("Initializing services...")

	// Initialize the PostgreSQL database
	dbPostgres := databases.InitializePostgres()
	l.InfoMsg("PostgreSQL database initialized successfully")

	// Initialize Redis client
	redisClient := redisdb.InitializeRedis()
	l.InfoMsg("Redis client initialized successfully")

	// Initialize InfluxDB client
	influxClient := influxdb.InitializeInfluxDB()
	l.InfoMsg("InfluxDB client initialized successfully")

	return dbPostgres, redisClient, influxClient
}
