package startup

import (
	"github.com/fatih/color"
	"github.com/go-redis/redis/v8"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	models "go-crud-api/internal/user/models"
	"go-crud-api/pkg/databases/influxdb"
	db "go-crud-api/pkg/databases/postgres"
	"go-crud-api/pkg/databases/redisdb"
	"gorm.io/gorm"
	"log"
)

// InitializeServices initializes all services and returns their instances.
func InitializeServices() (*gorm.DB, *redis.Client, influxdb2.Client) {
	log.Println("Initializing services...")

	// Initialize the PostgresSQL database using GORM
	dbPostgres := db.InitializePostgres()
	color.Blue("PostgresSQL database initialized successfully with GORM")

	// Initialize Redis client
	redisClient := redisdb.InitializeRedis()
	color.Green("Redis client initialized successfully")

	// Initialize InfluxDB client
	influxClient := influxdb.InitializeInfluxDB()
	color.Cyan("InfluxDB client initialized successfully")

	// Run database migrations
	MigrateDatabase(dbPostgres)

	return dbPostgres, redisClient, influxClient
}

// MigrateDatabase runs the database migrations and inserts test data if needed
func MigrateDatabase(db *gorm.DB) {
	log.Println("Checking and migrating database schema...")

	err := db.AutoMigrate(
		// Add models that need to be migrated
		&models.User{},
	)
	if err != nil {
		color.Red("Database migration failed: %v", err)
		return
	}

	color.Yellow("Database schema migrated successfully")
}

// InsertTestData inserts test data into the database
func InsertTestData(db *gorm.DB) {
	log.Println("Inserting test data...")

	user := models.User{
		Name:  "Test User",
		Email: "test@example.com",
		// Add more fields as necessary
	}

	result := db.Create(&user)
	if result.Error != nil {
		color.Red("Failed to insert test data: %v", result.Error)
	} else {
		color.Green("Test data inserted successfully")
	}
}
