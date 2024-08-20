package redisdb

import (
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

// InitializeRedis initializes the Redis client
func InitializeRedis() *redis.Client {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		log.Println("REDIS_URL environment variable is not set, falling back to default Redis URL")
		redisURL = "redis://localhost:6379/0" // Default to local Redis
	} else {
		log.Println("REDIS_URL environment variable loaded successfully")
	}

	options, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Fatalf("Failed to parse Redis URL: %v", err)
	}

	client := redis.NewClient(options)

	// Test connection
	ctx := context.Background()
	_, err = client.Ping(ctx).Result()
	if err != nil {
		log.Printf("Failed to connect to Redis at %s: %v\n", redisURL, err)
	} else {
		log.Printf("Successfully connected to Redis at %s\n", redisURL)
	}

	return client
}
