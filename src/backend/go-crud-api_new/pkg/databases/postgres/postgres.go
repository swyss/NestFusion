package postgres

import (
	task_models "go-crud-api/internal/tasks/models"
	models "go-crud-api/internal/user/models"
	"go-crud-api/utils"
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var PostgresDB *gorm.DB
var (
	database *gorm.DB
	once     sync.Once
)

// InitializePostgres initializes the PostgreSQL database connection using GORM.
func InitializePostgres() *gorm.DB {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		utils.PrintError("Error: DATABASE_URL environment variable is not set")
		log.Fatal("DATABASE_URL environment variable is not set")
	}
	utils.PrintSuccess("DATABASE_URL environment variable loaded successfully")

	var err error
	maxRetries := 10
	// Start the spinner while connecting to the database
	utils.StartSpinner(utils.FormatInfo, "Connecting to PostgreSQL")

	for i := 0; i < maxRetries; i++ {
		PostgresDB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err == nil {
			sqlDB, err := PostgresDB.DB()
			if err != nil {
				utils.PrintError("Failed to get database instance")
				log.Fatalf("Failed to get database instance: %v", err)
			}

			if err = sqlDB.Ping(); err == nil {
				utils.StopSpinner()
				utils.PrintSuccess("Successfully connected to PostgreSQL database using GORM")
				break
			}
		}

		utils.PrintWarning("Failed to ping the database (attempt %d/%d): %v", i+1, maxRetries, err)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		utils.StopSpinner()
		utils.PrintError("Could not connect to the database after %d attempts", maxRetries)
		log.Fatalf("Could not connect to the database after %d attempts: %v", maxRetries, err)
	}

	configureDBPooling(PostgresDB)

	if err := PostgresDB.AutoMigrate(&models.User{}); err == nil {
		utils.PrintSuccess("User table migration completed successfully")
	} else {
		utils.PrintError("Failed to migrate User table: %v", err)
		log.Fatalf("Failed to migrate User table: %v", err)
	}

	if err := PostgresDB.AutoMigrate(&task_models.Task{}); err == nil {
		utils.PrintSuccess("Task table migration completed successfully")
	} else {
		utils.PrintError("Failed to migrate Task table: %v", err)
		log.Fatalf("Failed to migrate Task table: %v", err)
	}
	return PostgresDB
}

// configureDBPooling configures the connection pool settings for the database.
func configureDBPooling(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		utils.PrintError("Failed to get database instance for pooling")
		log.Fatalf("Failed to get database instance: %v", err)
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	utils.PrintSuccess("Database connection pooling configured")
}

func initPostgres() {
	database = InitializePostgres()
}

func GetDBClient() *gorm.DB {
	once.Do(initPostgres)
	return database
}
