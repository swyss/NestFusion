package databases

import (
	"log"
	"os"
	"strconv"
	"time"

	usermodel "go-crud-api/internal/models/user"
	"gorm.io/driver/postgres" // PostgresSQL driver for GORM
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"go-crud-api/internal/models"
)

// InitializePostgres initializes the PostgresSQL database connection using GORM.
func InitializePostgres() *gorm.DB {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	} else {
		log.Println("DATABASE_URL environment variable loaded successfully")
	}

	var db *gorm.DB
	var err error
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // Uses singular table names, e.g., `User` instead of `Users`
			},
			Logger: logger.Default.LogMode(logger.Info), // Log SQL queries
		})
		if err == nil {
			sqlDB, err := db.DB()
			if err != nil {
				log.Fatalf("Failed to get database instance: %v", err)
			}

			// Test the connection to the database
			if err = sqlDB.Ping(); err == nil {
				log.Println("Successfully connected to PostgreSQL database using GORM")
				break
			}
		}

		log.Printf("Failed to ping the database (attempt %d/%d): %v\n", i+1, maxRetries, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("Could not connect to the database after %d attempts: %v", maxRetries, err)
	}

	// Configure connection pooling
	configureDBPooling(db)

	// AutoMigrate automatically migrates your schema, to keep your database schema up to date.
	if err := db.AutoMigrate(
		&usermodel.User{},
	); err != nil {
		log.Fatalf("Failed to migrate tables: %v", err)
	} else {
		log.Println("Database migration completed successfully")
	}

	if err := db.AutoMigrate(&models.Task{}); err == nil {
		log.Println("Database migration completed successfully")
	} else {
		log.Fatalf("Failed to migrate tables: %v", err)
	}
	return db
}

// configureDBPooling configures the connection pool settings for the database.
func configureDBPooling(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	// Set pool settings, possibly configurable via environment variables
	sqlDB.SetMaxOpenConns(getEnvAsInt("DB_MAX_OPEN_CONNS", 25))
	sqlDB.SetMaxIdleConns(getEnvAsInt("DB_MAX_IDLE_CONNS", 25))
	sqlDB.SetConnMaxLifetime(getEnvAsDuration("DB_CONN_MAX_LIFETIME", 5*time.Minute))
	log.Println("Database connection pooling configured")
}

// getEnvAsInt retrieves an environment variable as an integer, with a fallback default.
func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// getEnvAsDuration retrieves an environment variable as a time.Duration, with a fallback default.
func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	if value, exists := os.LookupEnv(key); exists {
		if durationValue, err := time.ParseDuration(value); err == nil {
			return durationValue
		}
	}
	return defaultValue
}
