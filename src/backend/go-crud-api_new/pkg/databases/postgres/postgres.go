package postgres

import (
	models "go-crud-api/internal/user/models"
	"gorm.io/driver/postgres" // PostgresSQL driver for GORM
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var PostgresDB *gorm.DB

// InitializePostgres initializes the PostgresSQL database connection using GORM.
func InitializePostgres() *gorm.DB {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	} else {
		log.Println("DATABASE_URL environment variable loaded successfully")
	}

	var err error
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		PostgresDB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{ // Correct use of gorm.Open with the Postgres driver
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // Uses singular table names, e.g., `User` instead of `Users`
			},
			Logger: logger.Default.LogMode(logger.Info), // Log SQL queries
		})
		if err == nil {
			sqlDB, err := PostgresDB.DB()
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
	configureDBPooling(PostgresDB)

	// AutoMigrate automatically migrates your schema, to keep your database schema up to date.
	if err := PostgresDB.AutoMigrate(&models.User{}); err == nil {
		log.Println("Database migration completed successfully")
	} else {
		log.Fatalf("Failed to migrate tables: %v", err)
	}

	return PostgresDB
}

// configureDBPooling configures the connection pool settings for the database.
func configureDBPooling(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	log.Println("Database connection pooling configured")
}
