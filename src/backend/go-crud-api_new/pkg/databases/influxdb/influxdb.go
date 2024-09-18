package influxdb

import (
	"context"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"go-crud-api/utils"
	"log"
	"os"
)

var InfluxClient influxdb2.Client

// InitializeInfluxDB initializes the InfluxDB client
func InitializeInfluxDB() influxdb2.Client {
	influxURL := os.Getenv("INFLUXDB_URL")
	if influxURL == "" {
		utils.PrintError("Error: INFLUXDB_URL environment variable is not set")
		log.Fatal("INFLUXDB_URL environment variable is not set")
	}
	utils.PrintSuccess("INFLUXDB_URL environment variable loaded successfully")

	token := os.Getenv("INFLUXDB_TOKEN")
	if token == "" {
		utils.PrintError("Error: INFLUXDB_TOKEN environment variable is not set")
		log.Fatal("INFLUXDB_TOKEN environment variable is not set")
	}
	utils.PrintSuccess("INFLUXDB_TOKEN environment variable loaded successfully")

	InfluxClient := influxdb2.NewClient(influxURL, token)

	// Use a spinner while trying to connect to InfluxDB
	utils.StartSpinner(utils.FormatInfo, "Connecting to InfluxDB")

	// Ping InfluxDB to check if it is reachable
	_, err := InfluxClient.Ready(context.Background())
	if err != nil {
		utils.StopSpinner()
		utils.PrintError("Failed to connect to InfluxDB")
		log.Fatalf("Failed to connect to InfluxDB: %v", err)
	} else {
		utils.StopSpinner()
		utils.PrintSuccess("Successfully connected to InfluxDB")
	}

	return InfluxClient
}
