package startup

import (
	"github.com/fatih/color"
	"github.com/go-redis/redis/v8"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"go-crud-api/internal/user/models"
	"go-crud-api/pkg/databases/influxdb"
	"go-crud-api/pkg/databases/postgres"
	"go-crud-api/pkg/databases/redisdb"
	"go-crud-api/utils"
	"gorm.io/gorm"
	"log"
	"time"
)

// Wait for PostgreSQL to be ready before establishing a connection
func waitForPostgres() *gorm.DB {
	retries := 5
	var dbPostgres *gorm.DB
	var err error
	for i := 0; i < retries; i++ {
		dbPostgres, err = postgres.InitializePostgres()
		if err == nil {
			utils.PrintSuccess("Connected to PostgreSQL")
			return dbPostgres
		}
		log.Printf("PostgreSQL not ready yet, retrying in 5 seconds... (attempt %d/%d)", i+1, retries)
		time.Sleep(5 * time.Second)
	}
	log.Fatal("Could not connect to PostgreSQL after multiple attempts")
	return nil
}

// InitializeServices initializes all services and returns their instances
func InitializeServices() (*gorm.DB, *redis.Client, influxdb2.Client) {
	utils.PrintInfo("Initializing services...")

	// Wait for PostgreSQL and establish the connection
	dbPostgres := waitForPostgres()
	if dbPostgres == nil {
		color.Red("Error initializing PostgreSQL.")
		return nil, nil, nil
	} else {
		utils.PrintSuccess("PostgreSQL initialized successfully.")
	}

	// Perform database migrations
	err := postgres.MigrateDatabase(dbPostgres)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	utils.PrintSuccess("Database migrations completed.")

	// Initialize Redis client
	redisClient := redisdb.InitializeRedis()
	if redisClient == nil {
		color.Red("Error initializing Redis.")
	} else {
		utils.PrintSuccess("Redis initialized successfully.")
	}

	// Initialize InfluxDB client
	influxClient := influxdb.InitializeInfluxDB()
	if influxClient == nil {
		color.Red("Error initializing InfluxDB.")
	} else {
		utils.PrintSuccess("InfluxDB initialized successfully.")
	}

	// Seed the database with initial data
	SeedDatabase(dbPostgres)

	return dbPostgres, redisClient, influxClient
}

// SeedDatabase inserts test data into the PostgreSQL database
func SeedDatabase(db *gorm.DB) {
	log.Println("Seeding database with test data...")

	// Create roles
	adminRole := models.UserRole{RoleName: "Admin"}
	userRole := models.UserRole{RoleName: "User"}
	managerRole := models.UserRole{RoleName: "Manager"}

	// Insert roles into the database
	db.FirstOrCreate(&adminRole, models.UserRole{RoleName: "Admin"})
	db.FirstOrCreate(&userRole, models.UserRole{RoleName: "User"})
	db.FirstOrCreate(&managerRole, models.UserRole{RoleName: "Manager"})

	// Create users
	admin := models.User{
		UserName: "admin",
		Password: "adminpass", // This password should be hashed in a real application
		RoleID:   adminRole.ID,
	}
	user := models.User{
		UserName: "user",
		Password: "userpass", // This password should be hashed in a real application
		RoleID:   userRole.ID,
	}
	manager := models.User{
		UserName: "manager",
		Password: "managerpass", // This password should be hashed in a real application
		RoleID:   managerRole.ID,
	}

	// Insert users into the database
	db.FirstOrCreate(&admin, models.User{UserName: "admin"})
	db.FirstOrCreate(&user, models.User{UserName: "user"})
	db.FirstOrCreate(&manager, models.User{UserName: "manager"})

	// Create UserInfo for each user
	adminInfo := models.UserInfo{
		UserID: admin.ID,
		Name:   "Admin User",
		Email:  "admin@example.com",
	}
	userInfo := models.UserInfo{
		UserID: user.ID,
		Name:   "Regular User",
		Email:  "user@example.com",
	}
	managerInfo := models.UserInfo{
		UserID: manager.ID,
		Name:   "Manager User",
		Email:  "manager@example.com",
	}

	// Insert UserInfo into the database
	db.FirstOrCreate(&adminInfo, models.UserInfo{UserID: admin.ID})
	db.FirstOrCreate(&userInfo, models.UserInfo{UserID: user.ID})
	db.FirstOrCreate(&managerInfo, models.UserInfo{UserID: manager.ID})

	utils.PrintSuccess("Database seeded with test data successfully.")
}
