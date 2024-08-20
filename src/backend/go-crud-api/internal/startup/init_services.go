package startup

import (
	"github.com/go-redis/redis/v8"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"go-crud-api/internal/databases"
	"go-crud-api/internal/databases/influxdb"
	"go-crud-api/internal/databases/redisdb"
	"gorm.io/gorm"
	"log"
)

// InitializeServices initializes all services and returns their instances.
func InitializeServices() (*gorm.DB, *redis.Client, influxdb2.Client) {
	log.Println("Initializing services...")

	// Initialize the PostgresSQL database using GORM
	dbPostgres := databases.InitializePostgres()
	log.Println("PostgresSQL database initialized successfully with GORM")

	// Initialize Redis client
	redisClient := redisdb.InitializeRedis()
	log.Println("Redis client initialized successfully")

	// Initialize InfluxDB client
	influxClient := influxdb.InitializeInfluxDB()
	log.Println("InfluxDB client initialized successfully")

	return dbPostgres, redisClient, influxClient
}
