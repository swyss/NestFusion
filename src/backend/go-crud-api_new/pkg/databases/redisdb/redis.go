package redisdb

import (
	"go-crud-api/utils"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

var RedisClient *redis.Client
var ctx = context.Background()

// InitializeRedis initializes the Redis client
func InitializeRedis() *redis.Client {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		utils.PrintWarning("REDIS_URL environment variable is not set, falling back to default Redis URL")
		redisURL = "redis://localhost:6379/0" // Default to local Redis
	} else {
		utils.PrintSuccess("REDIS_URL environment variable loaded successfully")
	}

	options, err := redis.ParseURL(redisURL)
	if err != nil {
		utils.PrintError("Error: Failed to parse Redis URL")
		log.Fatalf("Failed to parse Redis URL: %v", err)
	}

	RedisClient := redis.NewClient(options)

	// Use a spinner while connecting to Redis
	utils.StartSpinner(utils.FormatInfo, "Connecting to Redis")

	// Test connection
	ctx := context.Background()
	_, err = RedisClient.Ping(ctx).Result()
	if err != nil {
		utils.StopSpinner()
		utils.PrintError("Failed to connect to Redis at %s", redisURL)
		log.Printf("Failed to connect to Redis at %s: %v\n", redisURL, err)
	} else {
		utils.StopSpinner()
		utils.PrintSuccess("Successfully connected to Redis at %s", redisURL)
	}

	return RedisClient
}
