package influxdb

import (
	"context"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"log"
	"os"
)

var InfluxClient influxdb2.Client

// InitializeInfluxDB initializes the InfluxDB client
func InitializeInfluxDB() influxdb2.Client {
	influxURL := os.Getenv("INFLUXDB_URL")
	if influxURL == "" {
		log.Fatal("INFLUXDB_URL environment variable is not set")
	}
	log.Println("INFLUXDB_URL environment variable loaded successfully")

	token := os.Getenv("INFLUXDB_TOKEN")
	if token == "" {
		log.Fatal("INFLUXDB_TOKEN environment variable is not set")
	}
	log.Println("INFLUXDB_TOKEN environment variable loaded successfully")

	InfluxClient := influxdb2.NewClient(influxURL, token)

	// Ping InfluxDB to check if it is reachable
	_, err := InfluxClient.Ready(context.Background())
	if err != nil {
		log.Fatalf("Failed to connect to InfluxDB: %v", err)
	} else {
		log.Println("Successfully connected to InfluxDB")
	}

	return InfluxClient
}
