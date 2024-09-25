package postgres

import (
	"fmt"
	task_models "go-crud-api/internal/tasks/models"
	user_models "go-crud-api/internal/user/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"sync"
)

var PostgresDB *gorm.DB
var (
	database *gorm.DB
	once     sync.Once
)

// InitializePostgres establishes a connection to the PostgreSQL database
func InitializePostgres() (*gorm.DB, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	fmt.Println(databaseURL)
	if databaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable not set")
	}

	// Connect to the PostgreSQL database
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	return db, nil
}

// MigrateDatabase applies migrations to the PostgreSQL database
func MigrateDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(&user_models.UserRole{}, &user_models.User{}, &task_models.Task{}, &user_models.UserInfo{})
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}
	return nil
}
