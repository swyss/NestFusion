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

	// Configure the GORM settings
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		utils.PrintError("Failed to connect to PostgreSQL")
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	utils.PrintSuccess("Connected to PostgreSQL")

	// Configure connection pooling
	configureDBPooling(db)

	// Migrate the schema
	db.AutoMigrate(&models.User{}, &task_models.Task{})
	PostgresDB = db

	return db
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
