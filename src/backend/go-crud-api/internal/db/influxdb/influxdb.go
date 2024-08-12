package influxdb

import (
	"context"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"log"
	"os"
)

// InitializeInfluxDB initializes the InfluxDB client
func InitializeInfluxDB() influxdb2.Client {
	influxURL := os.Getenv("INFLUXDB_URL")
	if influxURL == "" {
		log.Fatal("INFLUXDB_URL environment variable is not set")
	}

	token := os.Getenv("INFLUXDB_TOKEN")
	if token == "" {
		log.Fatal("INFLUXDB_TOKEN environment variable is not set")
	}

	client := influxdb2.NewClient(influxURL, token)

	// Ping InfluxDB to check if it is reachable
	_, err := client.Ready(context.Background())
	if err != nil {
		log.Fatalf("Failed to connect to InfluxDB: %v", err)
	}

	return client
}
