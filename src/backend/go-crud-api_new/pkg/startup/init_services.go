package startup

import (
	"github.com/fatih/color"
	"github.com/go-redis/redis/v8"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	models "go-crud-api/internal/user/models"
	"go-crud-api/pkg/databases/influxdb"
	postgres "go-crud-api/pkg/databases/postgres"
	redisdb "go-crud-api/pkg/databases/redisdb"

	"go-crud-api/utils"

	"gorm.io/gorm"
	"log"
)

// InitializeServices initializes all services and returns their instances.
func InitializeServices() (*gorm.DB, *redis.Client, influxdb2.Client) {
	utils.PrintInfo("Initializing services...")

	// Initialize the PostgreSQL database using GORM
	dbPostgres := postgres.InitializePostgres()
	if dbPostgres == nil {
		color.Red("Error initializing PostgreSQL.")
	} else {
		utils.PrintSuccess("PostgreSQL initialized successfully.")
	}

	// Initialize the Redis client
	redisClient := redisdb.InitializeRedis()
	if redisClient == nil {
		color.Red("Error initializing Redis.")
	} else {
		utils.PrintSuccess("Redis initialized successfully.")
	}

	// Initialize the InfluxDB client
	influxClient := influxdb.InitializeInfluxDB()
	if influxClient == nil {
		color.Red("Error initializing InfluxDB.")
	} else {
		utils.PrintSuccess("InfluxDB initialized successfully.")
	}

	// Seed the PostgreSQL database with sample data
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
